// Code generated by MockGen. DO NOT EDIT.
// Source: goms.io/aks/rp/core/retry (interfaces: SingleIterationInterface,RetryerInterface)

// Package mock_retry is a generated GoMock package.
package mock_retry

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	retry "github.com/Azure/aks-deployer/pkg/retry"
	reflect "reflect"
)

// MockSingleIterationInterface is a mock of SingleIterationInterface interface
type MockSingleIterationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSingleIterationInterfaceMockRecorder
}

// MockSingleIterationInterfaceMockRecorder is the mock recorder for MockSingleIterationInterface
type MockSingleIterationInterfaceMockRecorder struct {
	mock *MockSingleIterationInterface
}

// NewMockSingleIterationInterface creates a new mock instance
func NewMockSingleIterationInterface(ctrl *gomock.Controller) *MockSingleIterationInterface {
	mock := &MockSingleIterationInterface{ctrl: ctrl}
	mock.recorder = &MockSingleIterationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSingleIterationInterface) EXPECT() *MockSingleIterationInterfaceMockRecorder {
	return m.recorder
}

// RunOnce mocks base method
func (m *MockSingleIterationInterface) RunOnce(arg0 context.Context) (retry.Status, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOnce", arg0)
	ret0, _ := ret[0].(retry.Status)
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunOnce indicates an expected call of RunOnce
func (mr *MockSingleIterationInterfaceMockRecorder) RunOnce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOnce", reflect.TypeOf((*MockSingleIterationInterface)(nil).RunOnce), arg0)
}

// MockRetryerInterface is a mock of RetryerInterface interface
type MockRetryerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRetryerInterfaceMockRecorder
}

// MockRetryerInterfaceMockRecorder is the mock recorder for MockRetryerInterface
type MockRetryerInterfaceMockRecorder struct {
	mock *MockRetryerInterface
}

// NewMockRetryerInterface creates a new mock instance
func NewMockRetryerInterface(ctrl *gomock.Controller) *MockRetryerInterface {
	mock := &MockRetryerInterface{ctrl: ctrl}
	mock.recorder = &MockRetryerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRetryerInterface) EXPECT() *MockRetryerInterfaceMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRetryerInterface) Run(arg0 context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockRetryerInterfaceMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRetryerInterface)(nil).Run), arg0)
}