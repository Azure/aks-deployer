// Code generated by MockGen. DO NOT EDIT.
// Source: goms.io/aks/rp/core/clients/keyvaultsecretclient (interfaces: Interface)

// Package mock_keyvaultsecretclient is a generated GoMock package.
package mock_keyvaultsecretclient

import (
	context "context"
	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	retry "github.com/Azure/aks-deployer/pkg/retry"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method
func (m *MockInterface) DeleteSecret(arg0 context.Context, arg1, arg2 string) (*keyvault.DeletedSecretBundle, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(*keyvault.DeletedSecretBundle)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockInterfaceMockRecorder) DeleteSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockInterface)(nil).DeleteSecret), arg0, arg1, arg2)
}

// GetDeletedSecrets mocks base method
func (m *MockInterface) GetDeletedSecrets(arg0 context.Context, arg1 string, arg2 *int32) (*keyvault.DeletedSecretListResult, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletedSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(*keyvault.DeletedSecretListResult)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetDeletedSecrets indicates an expected call of GetDeletedSecrets
func (mr *MockInterfaceMockRecorder) GetDeletedSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletedSecrets", reflect.TypeOf((*MockInterface)(nil).GetDeletedSecrets), arg0, arg1, arg2)
}

// GetSecret mocks base method
func (m *MockInterface) GetSecret(arg0 context.Context, arg1, arg2, arg3 string) (*keyvault.SecretBundle, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*keyvault.SecretBundle)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockInterfaceMockRecorder) GetSecret(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockInterface)(nil).GetSecret), arg0, arg1, arg2, arg3)
}

// GetSecrets mocks base method
func (m *MockInterface) GetSecrets(arg0 context.Context, arg1 string, arg2 *int32) (*keyvault.SecretListResult, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(*keyvault.SecretListResult)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets
func (mr *MockInterfaceMockRecorder) GetSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockInterface)(nil).GetSecrets), arg0, arg1, arg2)
}

// PurgeDeletedSecret mocks base method
func (m *MockInterface) PurgeDeletedSecret(arg0 context.Context, arg1, arg2 string) (*autorest.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeDeletedSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(*autorest.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PurgeDeletedSecret indicates an expected call of PurgeDeletedSecret
func (mr *MockInterfaceMockRecorder) PurgeDeletedSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDeletedSecret", reflect.TypeOf((*MockInterface)(nil).PurgeDeletedSecret), arg0, arg1, arg2)
}

// SetSecret mocks base method
func (m *MockInterface) SetSecret(arg0 context.Context, arg1, arg2 string, arg3 keyvault.SecretSetParameters) (*keyvault.SecretBundle, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*keyvault.SecretBundle)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// SetSecret indicates an expected call of SetSecret
func (mr *MockInterfaceMockRecorder) SetSecret(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockInterface)(nil).SetSecret), arg0, arg1, arg2, arg3)
}