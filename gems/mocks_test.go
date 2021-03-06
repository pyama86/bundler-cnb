// Code generated by MockGen. DO NOT EDIT.
// Source: gems.go

// Package gems_test is a generated GoMock package.
package gems_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPackageManager is a mock of PackageManager interface
type MockPackageManager struct {
	ctrl     *gomock.Controller
	recorder *MockPackageManagerMockRecorder
}

// MockPackageManagerMockRecorder is the mock recorder for MockPackageManager
type MockPackageManagerMockRecorder struct {
	mock *MockPackageManager
}

// NewMockPackageManager creates a new mock instance
func NewMockPackageManager(ctrl *gomock.Controller) *MockPackageManager {
	mock := &MockPackageManager{ctrl: ctrl}
	mock.recorder = &MockPackageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPackageManager) EXPECT() *MockPackageManagerMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockPackageManager) Install(location string) error {
	ret := m.ctrl.Call(m, "Install", location)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockPackageManagerMockRecorder) Install(location interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockPackageManager)(nil).Install), location)
}
