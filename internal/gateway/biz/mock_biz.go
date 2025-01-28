// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onexstack/onex/internal/gateway/biz (interfaces: IBiz)

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	miner "github.com/onexstack/onex/internal/gateway/biz/v1/miner"
	minerset "github.com/onexstack/onex/internal/gateway/biz/v1/minerset"
)

// MockIBiz is a mock of IBiz interface.
type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

// MockIBizMockRecorder is the mock recorder for MockIBiz.
type MockIBizMockRecorder struct {
	mock *MockIBiz
}

// NewMockIBiz creates a new mock instance.
func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}

// MinerSetV1 mocks base method.
func (m *MockIBiz) MinerSetV1() minerset.MinerSetBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinerSetV1")
	ret0, _ := ret[0].(minerset.MinerSetBiz)
	return ret0
}

// MinerSetV1 indicates an expected call of MinerSetV1.
func (mr *MockIBizMockRecorder) MinerSetV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinerSetV1", reflect.TypeOf((*MockIBiz)(nil).MinerSetV1))
}

// MinerV1 mocks base method.
func (m *MockIBiz) MinerV1() miner.MinerBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinerV1")
	ret0, _ := ret[0].(miner.MinerBiz)
	return ret0
}

// MinerV1 indicates an expected call of MinerV1.
func (mr *MockIBizMockRecorder) MinerV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinerV1", reflect.TypeOf((*MockIBiz)(nil).MinerV1))
}
