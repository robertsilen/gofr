// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go
//
// Generated by this command:
//
//	mockgen -destination=mock_client.go -package=nats -source=./interfaces.go Client,Subscription,ConnInterface,ConnectionManagerInterface,SubscriptionManagerInterface,StreamManagerInterface
//

// Package nats is a generated GoMock package.
package nats

import (
	context "context"
	reflect "reflect"

	nats "github.com/nats-io/nats.go"
	jetstream "github.com/nats-io/nats.go/jetstream"
	gomock "go.uber.org/mock/gomock"
	datasource "gofr.dev/pkg/gofr/datasource"
	pubsub "gofr.dev/pkg/gofr/datasource/pubsub"
)

// MockConnInterface is a mock of ConnInterface interface.
type MockConnInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnInterfaceMockRecorder
}

// MockConnInterfaceMockRecorder is the mock recorder for MockConnInterface.
type MockConnInterfaceMockRecorder struct {
	mock *MockConnInterface
}

// NewMockConnInterface creates a new mock instance.
func NewMockConnInterface(ctrl *gomock.Controller) *MockConnInterface {
	mock := &MockConnInterface{ctrl: ctrl}
	mock.recorder = &MockConnInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnInterface) EXPECT() *MockConnInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnInterface) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockConnInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnInterface)(nil).Close))
}

// JetStream mocks base method.
func (m *MockConnInterface) JetStream() (jetstream.JetStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "jStream")
	ret0, _ := ret[0].(jetstream.JetStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JetStream indicates an expected call of JetStream.
func (mr *MockConnInterfaceMockRecorder) JetStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "jStream", reflect.TypeOf((*MockConnInterface)(nil).JetStream))
}

// NATSConn mocks base method.
func (m *MockConnInterface) NATSConn() *nats.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NATSConn")
	ret0, _ := ret[0].(*nats.Conn)
	return ret0
}

// NATSConn indicates an expected call of NATSConn.
func (mr *MockConnInterfaceMockRecorder) NATSConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NATSConn", reflect.TypeOf((*MockConnInterface)(nil).NATSConn))
}

// Status mocks base method.
func (m *MockConnInterface) Status() nats.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(nats.Status)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockConnInterfaceMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockConnInterface)(nil).Status))
}

// MockNATSConnector is a mock of Connector interface.
type MockNATSConnector struct {
	ctrl     *gomock.Controller
	recorder *MockNATSConnectorMockRecorder
}

// MockNATSConnectorMockRecorder is the mock recorder for MockNATSConnector.
type MockNATSConnectorMockRecorder struct {
	mock *MockNATSConnector
}

// NewMockNATSConnector creates a new mock instance.
func NewMockNATSConnector(ctrl *gomock.Controller) *MockNATSConnector {
	mock := &MockNATSConnector{ctrl: ctrl}
	mock.recorder = &MockNATSConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNATSConnector) EXPECT() *MockNATSConnectorMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockNATSConnector) Connect(arg0 string, arg1 ...nats.Option) (ConnInterface, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Connect", varargs...)
	ret0, _ := ret[0].(ConnInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockNATSConnectorMockRecorder) Connect(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNATSConnector)(nil).Connect), varargs...)
}

// MockJetStreamCreator is a mock of JetStreamCreator interface.
type MockJetStreamCreator struct {
	ctrl     *gomock.Controller
	recorder *MockJetStreamCreatorMockRecorder
}

// MockJetStreamCreatorMockRecorder is the mock recorder for MockJetStreamCreator.
type MockJetStreamCreatorMockRecorder struct {
	mock *MockJetStreamCreator
}

// NewMockJetStreamCreator creates a new mock instance.
func NewMockJetStreamCreator(ctrl *gomock.Controller) *MockJetStreamCreator {
	mock := &MockJetStreamCreator{ctrl: ctrl}
	mock.recorder = &MockJetStreamCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJetStreamCreator) EXPECT() *MockJetStreamCreatorMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockJetStreamCreator) New(conn ConnInterface) (jetstream.JetStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", conn)
	ret0, _ := ret[0].(jetstream.JetStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockJetStreamCreatorMockRecorder) New(conn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockJetStreamCreator)(nil).New), conn)
}

// MockJetStreamClient is a mock of JetStreamClient interface.
type MockJetStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockJetStreamClientMockRecorder
}

// MockJetStreamClientMockRecorder is the mock recorder for MockJetStreamClient.
type MockJetStreamClientMockRecorder struct {
	mock *MockJetStreamClient
}

// NewMockJetStreamClient creates a new mock instance.
func NewMockJetStreamClient(ctrl *gomock.Controller) *MockJetStreamClient {
	mock := &MockJetStreamClient{ctrl: ctrl}
	mock.recorder = &MockJetStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJetStreamClient) EXPECT() *MockJetStreamClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockJetStreamClient) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockJetStreamClientMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockJetStreamClient)(nil).Close), ctx)
}

// CreateOrUpdateStream mocks base method.
func (m *MockJetStreamClient) CreateOrUpdateStream(ctx context.Context, cfg jetstream.StreamConfig) (jetstream.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateStream", ctx, cfg)
	ret0, _ := ret[0].(jetstream.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateStream indicates an expected call of CreateOrUpdateStream.
func (mr *MockJetStreamClientMockRecorder) CreateOrUpdateStream(ctx, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateStream", reflect.TypeOf((*MockJetStreamClient)(nil).CreateOrUpdateStream), ctx, cfg)
}

// CreateStream mocks base method.
func (m *MockJetStreamClient) CreateStream(ctx context.Context, cfg StreamConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", ctx, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockJetStreamClientMockRecorder) CreateStream(ctx, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockJetStreamClient)(nil).CreateStream), ctx, cfg)
}

// DeleteStream mocks base method.
func (m *MockJetStreamClient) DeleteStream(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockJetStreamClientMockRecorder) DeleteStream(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockJetStreamClient)(nil).DeleteStream), ctx, name)
}

// Health mocks base method.
func (m *MockJetStreamClient) Health() datasource.Health {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(datasource.Health)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockJetStreamClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockJetStreamClient)(nil).Health))
}

// Publish mocks base method.
func (m *MockJetStreamClient) Publish(ctx context.Context, subject string, message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, subject, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockJetStreamClientMockRecorder) Publish(ctx, subject, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockJetStreamClient)(nil).Publish), ctx, subject, message)
}

// Subscribe mocks base method.
func (m *MockJetStreamClient) Subscribe(ctx context.Context, subject string, handler messageHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, subject, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockJetStreamClientMockRecorder) Subscribe(ctx, subject, handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockJetStreamClient)(nil).Subscribe), ctx, subject, handler)
}

// MockConnectionManagerInterface is a mock of ConnectionManagerInterface interface.
type MockConnectionManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionManagerInterfaceMockRecorder
}

// MockConnectionManagerInterfaceMockRecorder is the mock recorder for MockConnectionManagerInterface.
type MockConnectionManagerInterfaceMockRecorder struct {
	mock *MockConnectionManagerInterface
}

// NewMockConnectionManagerInterface creates a new mock instance.
func NewMockConnectionManagerInterface(ctrl *gomock.Controller) *MockConnectionManagerInterface {
	mock := &MockConnectionManagerInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionManagerInterface) EXPECT() *MockConnectionManagerInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnectionManagerInterface) Close(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", ctx)
}

// Close indicates an expected call of Close.
func (mr *MockConnectionManagerInterfaceMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnectionManagerInterface)(nil).Close), ctx)
}

// Connect mocks base method.
func (m *MockConnectionManagerInterface) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockConnectionManagerInterfaceMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConnectionManagerInterface)(nil).Connect))
}

// Health mocks base method.
func (m *MockConnectionManagerInterface) Health() datasource.Health {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(datasource.Health)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockConnectionManagerInterfaceMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockConnectionManagerInterface)(nil).Health))
}

// JetStream mocks base method.
func (m *MockConnectionManagerInterface) jetStream() (jetstream.JetStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "jStream")
	ret0, _ := ret[0].(jetstream.JetStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JetStream indicates an expected call of JetStream.
func (mr *MockConnectionManagerInterfaceMockRecorder) JetStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "jStream", reflect.TypeOf((*MockConnectionManagerInterface)(nil).jetStream))
}

// Publish mocks base method.
func (m *MockConnectionManagerInterface) Publish(ctx context.Context, subject string, message []byte, metrics Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, subject, message, metrics)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockConnectionManagerInterfaceMockRecorder) Publish(ctx, subject, message, metrics any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockConnectionManagerInterface)(nil).Publish), ctx, subject, message, metrics)
}

// MockSubscriptionManagerInterface is a mock of SubscriptionManagerInterface interface.
type MockSubscriptionManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionManagerInterfaceMockRecorder
}

// MockSubscriptionManagerInterfaceMockRecorder is the mock recorder for MockSubscriptionManagerInterface.
type MockSubscriptionManagerInterfaceMockRecorder struct {
	mock *MockSubscriptionManagerInterface
}

// NewMockSubscriptionManagerInterface creates a new mock instance.
func NewMockSubscriptionManagerInterface(ctrl *gomock.Controller) *MockSubscriptionManagerInterface {
	mock := &MockSubscriptionManagerInterface{ctrl: ctrl}
	mock.recorder = &MockSubscriptionManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionManagerInterface) EXPECT() *MockSubscriptionManagerInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSubscriptionManagerInterface) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSubscriptionManagerInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubscriptionManagerInterface)(nil).Close))
}

// Subscribe mocks base method.
func (m *MockSubscriptionManagerInterface) Subscribe(ctx context.Context, topic string, js jetstream.JetStream, cfg *Config, logger pubsub.Logger, metrics Metrics) (*pubsub.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, topic, js, cfg, logger, metrics)
	ret0, _ := ret[0].(*pubsub.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscriptionManagerInterfaceMockRecorder) Subscribe(ctx, topic, js, cfg, logger, metrics any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscriptionManagerInterface)(nil).Subscribe), ctx, topic, js, cfg, logger, metrics)
}

// MockStreamManagerInterface is a mock of StreamManagerInterface interface.
type MockStreamManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagerInterfaceMockRecorder
}

// MockStreamManagerInterfaceMockRecorder is the mock recorder for MockStreamManagerInterface.
type MockStreamManagerInterfaceMockRecorder struct {
	mock *MockStreamManagerInterface
}

// NewMockStreamManagerInterface creates a new mock instance.
func NewMockStreamManagerInterface(ctrl *gomock.Controller) *MockStreamManagerInterface {
	mock := &MockStreamManagerInterface{ctrl: ctrl}
	mock.recorder = &MockStreamManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManagerInterface) EXPECT() *MockStreamManagerInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdateStream mocks base method.
func (m *MockStreamManagerInterface) CreateOrUpdateStream(ctx context.Context, cfg *jetstream.StreamConfig) (jetstream.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateStream", ctx, cfg)
	ret0, _ := ret[0].(jetstream.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateStream indicates an expected call of CreateOrUpdateStream.
func (mr *MockStreamManagerInterfaceMockRecorder) CreateOrUpdateStream(ctx, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateStream", reflect.TypeOf((*MockStreamManagerInterface)(nil).CreateOrUpdateStream), ctx, cfg)
}

// CreateStream mocks base method.
func (m *MockStreamManagerInterface) CreateStream(ctx context.Context, cfg StreamConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", ctx, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockStreamManagerInterfaceMockRecorder) CreateStream(ctx, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockStreamManagerInterface)(nil).CreateStream), ctx, cfg)
}

// DeleteStream mocks base method.
func (m *MockStreamManagerInterface) DeleteStream(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamManagerInterfaceMockRecorder) DeleteStream(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamManagerInterface)(nil).DeleteStream), ctx, name)
}
