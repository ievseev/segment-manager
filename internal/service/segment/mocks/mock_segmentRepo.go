// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MocksegmentRepo is a mock of segmentRepo interface.
type MocksegmentRepo struct {
	ctrl     *gomock.Controller
	recorder *MocksegmentRepoMockRecorder
}

// MocksegmentRepoMockRecorder is the mock recorder for MocksegmentRepo.
type MocksegmentRepoMockRecorder struct {
	mock *MocksegmentRepo
}

// NewMocksegmentRepo creates a new mock instance.
func NewMocksegmentRepo(ctrl *gomock.Controller) *MocksegmentRepo {
	mock := &MocksegmentRepo{ctrl: ctrl}
	mock.recorder = &MocksegmentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksegmentRepo) EXPECT() *MocksegmentRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MocksegmentRepo) Delete(ctx context.Context, segmentName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, segmentName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MocksegmentRepoMockRecorder) Delete(ctx, segmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MocksegmentRepo)(nil).Delete), ctx, segmentName)
}

// Save mocks base method.
func (m *MocksegmentRepo) Save(ctx context.Context, segmentName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, segmentName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MocksegmentRepoMockRecorder) Save(ctx, segmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MocksegmentRepo)(nil).Save), ctx, segmentName)
}
