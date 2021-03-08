// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/user.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// String mocks base method.
func (m *MockUser) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockUserMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockUser)(nil).String))
}

// String2 mocks base method.
func (m *MockUser) String2() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String2")
	ret0, _ := ret[0].(string)
	return ret0
}

// String2 indicates an expected call of String2.
func (mr *MockUserMockRecorder) String2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String2", reflect.TypeOf((*MockUser)(nil).String2))
}
