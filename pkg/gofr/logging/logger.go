package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/term"

	"gofr.dev/pkg/gofr/version"
)

const (
	fileMode      = 0644
	maxBufferSize = 256 * 1024 //64kB
)

type PrettyPrint interface {
	PrettyPrint(writer io.Writer)
}

// Logger represents a logging interface.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	ChangeLevel(level Level)
	ShutDown()
}

type logger struct {
	level      Level
	out        io.Writer
	writer     *bufio.Writer
	ticker     *time.Ticker
	isTerminal bool
	lock       *sync.Mutex
	done       chan struct{}
}

type logEntry struct {
	Level       Level       `json:"level"`
	Time        time.Time   `json:"time"`
	Message     interface{} `json:"message"`
	GofrVersion string      `json:"gofrVersion"`
}

func (l *logger) Debug(args ...interface{}) {
	l.logf(DEBUG, "", args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logf(INFO, "", args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *logger) Notice(args ...interface{}) {
	l.logf(NOTICE, "", args...)
}

func (l *logger) Noticef(format string, args ...interface{}) {
	l.logf(NOTICE, format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logf(WARN, "", args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logf(WARN, format, args...)
}

func (l *logger) Log(args ...interface{}) {
	l.logf(INFO, "", args...)
}

func (l *logger) Logf(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logf(ERROR, "", args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logf(FATAL, "", args...)

	//nolint:revive // exit status is 1 as it denotes failure as signified by Fatal log
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logf(FATAL, format, args...)

	//nolint:revive // exit status is 1 as it denotes failure as signified by Fatal log
	os.Exit(1)
}

func (l *logger) ShutDown() {
	l.ticker.Stop()
	l.done <- struct{}{}
}

func (l *logger) logf(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	entry := logEntry{
		Level:       level,
		Time:        time.Now(),
		GofrVersion: version.Framework,
	}

	switch {
	case len(args) == 1 && format == "":
		entry.Message = args[0]
	case len(args) != 1 && format == "":
		entry.Message = args
	default:
		entry.Message = fmt.Sprintf(format, args...)
	}

	var bs []byte

	if l.isTerminal {
		bs = l.prettyPrint(entry)
	} else {
		bs, _ = json.Marshal(entry)
	}

	if len(bs) > l.writer.Available() && l.writer.Buffered() > 0 {
		if err := l.writer.Flush(); err != nil {
			// TODO: handle cases when the flush does not work
		}
	}

	l.writer.Write(bs)
}

func (l *logger) prettyPrint(e logEntry) []byte {
	// Note: we need to lock the pretty print as printing to stdandard output not concurency safe
	// the logs when printed in go routines were getting missaligned since we are achieveing
	// a single line of print, in 2 separate statements which caused the missalignment.
	l.lock.Lock()
	defer l.lock.Unlock()

	out := &bytes.Buffer{}

	// Pretty printing if the message interface defines a method PrettyPrint else print the print message
	// This decouples the logger implementation from its usage
	if fn, ok := e.Message.(PrettyPrint); ok {
		out.WriteString(fmt.Sprintf("\u001B[38;5;%dm%s\u001B[0m [%s] ", e.Level.color(), e.Level.String()[0:4],
			e.Time.Format("15:04:05")))

		fn.PrettyPrint(out)
	} else {
		out.WriteString(fmt.Sprintf("\u001B[38;5;%dm%s\u001B[0m [%s] %v\n", e.Level.color(), e.Level.String()[0:4],
			e.Time.Format("15:04:05"), e.Message))
	}

	return out.Bytes()
}

func (l *logger) startFlushLoop() {
	defer close(l.done)

	for {
		select {
		case <-l.ticker.C:
			_ = l.flush()
		case <-l.done:
			return
		}
	}
}

func (l *logger) flush() error {
	return l.writer.Flush()
}

// NewLogger creates a new logger instance with the specified logging level.
func NewLogger(level Level) Logger {
	l := &logger{
		writer: bufio.NewWriterSize(os.Stdout, maxBufferSize),
		ticker: time.NewTicker(1 * time.Second),
		lock:   new(sync.Mutex),
		done:   make(chan struct{}),
	}

	l.level = level
	l.isTerminal = checkIfTerminal(os.Stdout)

	go l.startFlushLoop()

	return l
}

// NewFileLogger creates a new logger instance with logging to a file.
func NewFileLogger(path string) Logger {
	l := &logger{out: io.Discard}

	if path == "" {
		return l
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		return l
	}

	l.out = f

	return l
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return term.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func (l *logger) ChangeLevel(level Level) {
	l.level = level
}
