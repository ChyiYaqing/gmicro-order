// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen -destination db_mock.go -package ports -source db.go
//

// Package ports is a generated GoMock package.
package ports

import (
	context "context"
	reflect "reflect"

	domain "github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockDBPort is a mock of DBPort interface.
type MockDBPort struct {
	ctrl     *gomock.Controller
	recorder *MockDBPortMockRecorder
	isgomock struct{}
}

// MockDBPortMockRecorder is the mock recorder for MockDBPort.
type MockDBPortMockRecorder struct {
	mock *MockDBPort
}

// NewMockDBPort creates a new mock instance.
func NewMockDBPort(ctrl *gomock.Controller) *MockDBPort {
	mock := &MockDBPort{ctrl: ctrl}
	mock.recorder = &MockDBPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBPort) EXPECT() *MockDBPortMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDBPort) Get(ctx context.Context, id int64) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDBPortMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBPort)(nil).Get), ctx, id)
}

// Save mocks base method.
func (m *MockDBPort) Save(arg0 context.Context, arg1 *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockDBPortMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDBPort)(nil).Save), arg0, arg1)
}
