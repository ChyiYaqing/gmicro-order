// Code generated by MockGen. DO NOT EDIT.
// Source: api.go
//
// Generated by this command:
//
//	mockgen -destination api_mock.go -package ports -source api.go
//

// Package ports is a generated GoMock package.
package ports

import (
	context "context"
	reflect "reflect"

	domain "github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockAPIPort is a mock of APIPort interface.
type MockAPIPort struct {
	ctrl     *gomock.Controller
	recorder *MockAPIPortMockRecorder
	isgomock struct{}
}

// MockAPIPortMockRecorder is the mock recorder for MockAPIPort.
type MockAPIPortMockRecorder struct {
	mock *MockAPIPort
}

// NewMockAPIPort creates a new mock instance.
func NewMockAPIPort(ctrl *gomock.Controller) *MockAPIPort {
	mock := &MockAPIPort{ctrl: ctrl}
	mock.recorder = &MockAPIPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIPort) EXPECT() *MockAPIPortMockRecorder {
	return m.recorder
}

// GetOrder mocks base method.
func (m *MockAPIPort) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, id)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockAPIPortMockRecorder) GetOrder(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockAPIPort)(nil).GetOrder), ctx, id)
}

// SaveOrder mocks base method.
func (m *MockAPIPort) SaveOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrder", ctx, order)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveOrder indicates an expected call of SaveOrder.
func (mr *MockAPIPortMockRecorder) SaveOrder(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrder", reflect.TypeOf((*MockAPIPort)(nil).SaveOrder), ctx, order)
}
