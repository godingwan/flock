// Code generated by MockGen. DO NOT EDIT.
// Source: flock.go

// Package flock is a generated GoMock package.
package flock

import (
	context "context"
	twitter "github.com/dghubble/go-twitter/twitter"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetTweets mocks base method
func (m *MockService) GetTweets(ctx context.Context, searchTerm string) (*twitter.Search, error) {
	ret := m.ctrl.Call(m, "GetTweets", ctx, searchTerm)
	ret0, _ := ret[0].(*twitter.Search)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTweets indicates an expected call of GetTweets
func (mr *MockServiceMockRecorder) GetTweets(ctx, searchTerm interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTweets", reflect.TypeOf((*MockService)(nil).GetTweets), ctx, searchTerm)
}
