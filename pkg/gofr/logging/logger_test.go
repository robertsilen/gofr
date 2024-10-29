package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/term"

	"gofr.dev/pkg/gofr/testutil"
)

func getLogger(t *testing.T, level Level) (*logger, *bytes.Buffer) {
	t.Helper()

	buf := &bytes.Buffer{}
	l := &logger{
		writer: bufio.NewWriterSize(buf, maxBufferSize),
		ticker: time.NewTicker(1 * time.Microsecond),
		lock:   new(sync.Mutex),
		rwLock: new(sync.Mutex),
		done:   make(chan struct{}),
		level:  level,
	}

	return l, buf
}

func TestLogger_LevelInfo(t *testing.T) {
	l, buf := getLogger(t, INFO)

	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Log("log")
	l.Error("error")

	l.flush()
	out := buf.String()

	assert.NotContains(t, out, "debug")
	assert.Contains(t, out, "info")
	assert.Contains(t, out, "log")
	assert.Contains(t, out, "notice")
	assert.Contains(t, out, "warn")
	assert.Contains(t, out, "error")
}

func TestLogger_LevelError(t *testing.T) {
	l, buf := getLogger(t, ERROR)

	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Log("log")
	l.Error("error")

	l.flush()
	out := buf.String()

	assert.NotContains(t, out, "debug")
	assert.NotContains(t, out, "info")
	assert.NotContains(t, out, "log")
	assert.NotContains(t, out, "notice")
	assert.NotContains(t, out, "warn")
	assert.Contains(t, out, "error")
}

func TestLogger_LevelDebug(t *testing.T) {
	l, buf := getLogger(t, DEBUG)

	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Log("log")
	l.Error("error")

	l.flush()
	out := buf.String()

	assert.Contains(t, out, "debug")
	assert.Contains(t, out, "info")
	assert.Contains(t, out, "notice")
	assert.Contains(t, out, "warn")
	assert.Contains(t, out, "log")
	assert.Contains(t, out, "error")
}

func TestLogger_LevelNotice(t *testing.T) {
	l, buf := getLogger(t, NOTICE)

	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Log("log")
	l.Error("error")

	l.flush()
	out := buf.String()

	assert.NotContains(t, out, "debug")
	assert.NotContains(t, out, "info")
	assert.NotContains(t, out, "log")
	assert.Contains(t, out, "notice")
	assert.Contains(t, out, "warn")
	assert.Contains(t, out, "error")
}

func TestLogger_LevelWarn(t *testing.T) {
	l, buf := getLogger(t, WARN)

	l.Debug("debug")
	l.Info("info")
	l.Log("log")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")

	l.flush()
	out := buf.String()

	assert.NotContains(t, out, "debug")
	assert.NotContains(t, out, "info")
	assert.NotContains(t, out, "log")
	assert.NotContains(t, out, "notice")
	assert.Contains(t, out, "warn")
	assert.Contains(t, out, "error")
}

//func TestLogger_LevelFatal(t *testing.T) {
//	// running the failing part only when a specific env variable is set
//	if os.Getenv("GOFR_EXITER") == "1" {
//		l, buf := getLogger(t, ERROR)
//
//		l.Debug("debug")
//		l.Info("info")
//		l.Notice("notice")
//		l.Warn("warn")
//		l.Log("log")
//		l.Error("error")
//
//		return
//	}
//
//	//nolint:gosec // starting the actual test in a different subprocess
//	cmd := exec.Command(os.Args[0], "-test.run=TestLogger_LevelFatal")
//	cmd.Env = append(os.Environ(), "GOFR_EXITER=1")
//
//	stdout, err := cmd.StderrPipe()
//	require.NoError(t, err)
//
//	require.NoError(t, cmd.Start())
//
//	logBytes, err := io.ReadAll(stdout)
//	require.NoError(t, err)
//
//	log := string(logBytes)
//
//	levels := []Level{DEBUG, INFO, NOTICE, WARN, ERROR} // levels which should not be present in case of FATAL log_level
//
//	for i, l := range levels {
//		assert.NotContainsf(t, log, l.String(), "TEST[%d], Failed.\nunexpected %s log", i, l)
//	}
//
//	assertMessageInJSONLog(t, log, "Test Fatal Log")
//
//	// Check that the program exited
//	err = cmd.Wait()
//
//	var e *exec.ExitError
//
//	require.ErrorAs(t, err, &e)
//	assert.False(t, e.Success())
//}

func assertMessageInJSONLog(t *testing.T, logLine, expectation string) {
	t.Helper()

	var l logEntry
	_ = json.Unmarshal([]byte(logLine), &l)

	if l.Message != expectation {
		t.Errorf("Log mismatch. Expected: %s Got: %s", expectation, l.Message)
	}
}

func TestCheckIfTerminal(t *testing.T) {
	tests := []struct {
		desc       string
		writer     io.Writer
		isTerminal bool
	}{
		{"Terminal Writer", os.Stdout, term.IsTerminal(int(os.Stdout.Fd()))},
		{"Non-Terminal Writer", os.Stderr, term.IsTerminal(int(os.Stderr.Fd()))},
		{"Non-Terminal Writer (not *os.File)", &bytes.Buffer{}, false},
	}

	for i, tc := range tests {
		result := checkIfTerminal(tc.writer)

		assert.Equal(t, tc.isTerminal, result, "TEST[%d], Failed.\n%s", i, tc.desc)
	}
}

func Test_NewSilentLoggerSTDOutput(t *testing.T) {
	logs := testutil.StdoutOutputForFunc(func() {
		l := NewFileLogger("")

		l.Info("Info Logs")
		l.Debug("Debug Logs")
		l.Notice("Notic Logs")
		l.Warn("Warn Logs")
		l.Infof("%v Logs", "Infof")
		l.Debugf("%v Logs", "Debugf")
		l.Noticef("%v Logs", "Noticef")
		l.Warnf("%v Logs", "warnf")
	})

	assert.Equal(t, "", logs)
}

type mockLog struct {
	msg string
}

func (m *mockLog) PrettyPrint(writer io.Writer) {
	fmt.Fprintf(writer, "TEST "+m.msg)
}

func TestPrettyPrint(t *testing.T) {
	m := &mockLog{msg: "mock test log"}
	out := &bytes.Buffer{}
	l := &logger{isTerminal: true, lock: new(sync.Mutex), rwLock: new(sync.Mutex)}

	// case PrettyPrint is implemented
	l.prettyPrint(logEntry{
		Level:   INFO,
		Message: m,
	})

	outputLog := out.String()
	expOut := []string{"INFO", "[00:00:00]", "TEST mock test log"}

	for _, v := range expOut {
		assert.Contains(t, outputLog, v)
	}

	// case pretty print is not implemented
	out.Reset()

	l.prettyPrint(logEntry{
		Level:   DEBUG,
		Message: "test log for normal log",
	})

	outputLog = out.String()
	expOut = []string{"DEBU", "[00:00:00]", "test log for normal log"}

	for _, v := range expOut {
		assert.Contains(t, outputLog, v)
	}
}

func BenchmarkLogger(b *testing.B) {
	lg := NewLogger(DEBUG)

	b.SetParallelism(100)
	// Set the number of parallel goroutines to use
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Infof("my new log")
			// Uncomment as needed for other log levels
			// lg.Info("my info log")
			// lg.Debug("my debug log")
			// lg.Error("my error log")
			// lg.Notice("my notice log")
			// lg.Warn("my warn log")
		}
	})
}

// 1000000             10236 ns/op             240 B/op          8 allocs/op
