// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/subnet (interfaces: Manager,KubeManager)

// Package mock_subnet is a generated GoMock package.
package mock_subnet

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	gomock "github.com/golang/mock/gomock"

	subnet "github.com/Azure/ARO-RP/pkg/util/subnet"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockManager) CreateOrUpdate(arg0 context.Context, arg1 string, arg2 *network.Subnet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockManagerMockRecorder) CreateOrUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockManager)(nil).CreateOrUpdate), arg0, arg1, arg2)
}

// CreateOrUpdateFromIds mocks base method.
func (m *MockManager) CreateOrUpdateFromIds(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateFromIds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateFromIds indicates an expected call of CreateOrUpdateFromIds.
func (mr *MockManagerMockRecorder) CreateOrUpdateFromIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateFromIds", reflect.TypeOf((*MockManager)(nil).CreateOrUpdateFromIds), arg0, arg1)
}

// Get mocks base method.
func (m *MockManager) Get(arg0 context.Context, arg1 string) (*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockManagerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManager)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockManager) GetAll(arg0 context.Context, arg1 []string) ([]*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockManagerMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockManager)(nil).GetAll), arg0, arg1)
}

// GetHighestFreeIP mocks base method.
func (m *MockManager) GetHighestFreeIP(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestFreeIP", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighestFreeIP indicates an expected call of GetHighestFreeIP.
func (mr *MockManagerMockRecorder) GetHighestFreeIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestFreeIP", reflect.TypeOf((*MockManager)(nil).GetHighestFreeIP), arg0, arg1)
}

// MockKubeManager is a mock of KubeManager interface.
type MockKubeManager struct {
	ctrl     *gomock.Controller
	recorder *MockKubeManagerMockRecorder
}

// MockKubeManagerMockRecorder is the mock recorder for MockKubeManager.
type MockKubeManagerMockRecorder struct {
	mock *MockKubeManager
}

// NewMockKubeManager creates a new mock instance.
func NewMockKubeManager(ctrl *gomock.Controller) *MockKubeManager {
	mock := &MockKubeManager{ctrl: ctrl}
	mock.recorder = &MockKubeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeManager) EXPECT() *MockKubeManagerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockKubeManager) List(arg0 context.Context) ([]subnet.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]subnet.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockKubeManagerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubeManager)(nil).List), arg0)
}
