// Code generated by MockGen. DO NOT EDIT.
// Source: domain/bet/bet.interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	challenged "github.com/TaApostado/bets/domain/challenged"
	gambler "github.com/TaApostado/bets/domain/gambler"
	notification "github.com/TaApostado/common/notification"
	gomock "github.com/golang/mock/gomock"
)

// MockIState is a mock of IState interface.
type MockIState struct {
	ctrl     *gomock.Controller
	recorder *MockIStateMockRecorder
}

// MockIStateMockRecorder is the mock recorder for MockIState.
type MockIStateMockRecorder struct {
	mock *MockIState
}

// NewMockIState creates a new mock instance.
func NewMockIState(ctrl *gomock.Controller) *MockIState {
	mock := &MockIState{ctrl: ctrl}
	mock.recorder = &MockIStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIState) EXPECT() *MockIStateMockRecorder {
	return m.recorder
}

// ChangeDescription mocks base method.
func (m *MockIState) ChangeDescription(description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDescription", description)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeDescription indicates an expected call of ChangeDescription.
func (mr *MockIStateMockRecorder) ChangeDescription(description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDescription", reflect.TypeOf((*MockIState)(nil).ChangeDescription), description)
}

// ChangeName mocks base method.
func (m *MockIState) ChangeName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeName indicates an expected call of ChangeName.
func (mr *MockIStateMockRecorder) ChangeName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockIState)(nil).ChangeName), name)
}

// ChangeValue mocks base method.
func (m *MockIState) ChangeValue(value float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeValue", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeValue indicates an expected call of ChangeValue.
func (mr *MockIStateMockRecorder) ChangeValue(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeValue", reflect.TypeOf((*MockIState)(nil).ChangeValue), value)
}

// Deposit mocks base method.
func (m *MockIState) Deposit(arg0 float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deposit indicates an expected call of Deposit.
func (mr *MockIStateMockRecorder) Deposit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockIState)(nil).Deposit), arg0)
}

// Withdraw mocks base method.
func (m *MockIState) Withdraw(arg0 float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockIStateMockRecorder) Withdraw(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockIState)(nil).Withdraw), arg0)
}

// MockIBet is a mock of IBet interface.
type MockIBet struct {
	ctrl     *gomock.Controller
	recorder *MockIBetMockRecorder
}

// MockIBetMockRecorder is the mock recorder for MockIBet.
type MockIBetMockRecorder struct {
	mock *MockIBet
}

// NewMockIBet creates a new mock instance.
func NewMockIBet(ctrl *gomock.Controller) *MockIBet {
	mock := &MockIBet{ctrl: ctrl}
	mock.recorder = &MockIBetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBet) EXPECT() *MockIBetMockRecorder {
	return m.recorder
}

// Challenged mocks base method.
func (m *MockIBet) Challenged() challenged.IChallenged {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Challenged")
	ret0, _ := ret[0].(challenged.IChallenged)
	return ret0
}

// Challenged indicates an expected call of Challenged.
func (mr *MockIBetMockRecorder) Challenged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Challenged", reflect.TypeOf((*MockIBet)(nil).Challenged))
}

// ChangeDescription mocks base method.
func (m *MockIBet) ChangeDescription(description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDescription", description)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeDescription indicates an expected call of ChangeDescription.
func (mr *MockIBetMockRecorder) ChangeDescription(description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDescription", reflect.TypeOf((*MockIBet)(nil).ChangeDescription), description)
}

// ChangeName mocks base method.
func (m *MockIBet) ChangeName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeName indicates an expected call of ChangeName.
func (mr *MockIBetMockRecorder) ChangeName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockIBet)(nil).ChangeName), name)
}

// ChangeValue mocks base method.
func (m *MockIBet) ChangeValue(value float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeValue", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeValue indicates an expected call of ChangeValue.
func (mr *MockIBetMockRecorder) ChangeValue(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeValue", reflect.TypeOf((*MockIBet)(nil).ChangeValue), value)
}

// Close mocks base method.
func (m *MockIBet) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIBetMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIBet)(nil).Close))
}

// Credits mocks base method.
func (m *MockIBet) Credits() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credits")
	ret0, _ := ret[0].(float32)
	return ret0
}

// Credits indicates an expected call of Credits.
func (mr *MockIBetMockRecorder) Credits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credits", reflect.TypeOf((*MockIBet)(nil).Credits))
}

// Deposit mocks base method.
func (m *MockIBet) Deposit(arg0 float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deposit indicates an expected call of Deposit.
func (mr *MockIBetMockRecorder) Deposit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockIBet)(nil).Deposit), arg0)
}

// Description mocks base method.
func (m *MockIBet) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockIBetMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockIBet)(nil).Description))
}

// Gamblers mocks base method.
func (m *MockIBet) Gamblers() []gambler.IGambler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gamblers")
	ret0, _ := ret[0].([]gambler.IGambler)
	return ret0
}

// Gamblers indicates an expected call of Gamblers.
func (mr *MockIBetMockRecorder) Gamblers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gamblers", reflect.TypeOf((*MockIBet)(nil).Gamblers))
}

// Id mocks base method.
func (m *MockIBet) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockIBetMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockIBet)(nil).Id))
}

// IsClose mocks base method.
func (m *MockIBet) IsClose() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClose")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClose indicates an expected call of IsClose.
func (mr *MockIBetMockRecorder) IsClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClose", reflect.TypeOf((*MockIBet)(nil).IsClose))
}

// IsOpen mocks base method.
func (m *MockIBet) IsOpen() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpen")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOpen indicates an expected call of IsOpen.
func (mr *MockIBetMockRecorder) IsOpen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpen", reflect.TypeOf((*MockIBet)(nil).IsOpen))
}

// IsSuspend mocks base method.
func (m *MockIBet) IsSuspend() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSuspend")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSuspend indicates an expected call of IsSuspend.
func (mr *MockIBetMockRecorder) IsSuspend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSuspend", reflect.TypeOf((*MockIBet)(nil).IsSuspend))
}

// Name mocks base method.
func (m *MockIBet) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockIBetMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockIBet)(nil).Name))
}

// Notificator mocks base method.
func (m *MockIBet) Notificator() notification.INotificator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notificator")
	ret0, _ := ret[0].(notification.INotificator)
	return ret0
}

// Notificator indicates an expected call of Notificator.
func (mr *MockIBetMockRecorder) Notificator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notificator", reflect.TypeOf((*MockIBet)(nil).Notificator))
}

// Open mocks base method.
func (m *MockIBet) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockIBetMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockIBet)(nil).Open))
}

// Suspend mocks base method.
func (m *MockIBet) Suspend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend")
	ret0, _ := ret[0].(error)
	return ret0
}

// Suspend indicates an expected call of Suspend.
func (mr *MockIBetMockRecorder) Suspend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockIBet)(nil).Suspend))
}

// Value mocks base method.
func (m *MockIBet) Value() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(float32)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockIBetMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockIBet)(nil).Value))
}

// Withdraw mocks base method.
func (m *MockIBet) Withdraw(arg0 float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockIBetMockRecorder) Withdraw(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockIBet)(nil).Withdraw), arg0)
}
