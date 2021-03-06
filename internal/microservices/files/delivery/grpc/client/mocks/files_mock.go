// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/microservices/files/delivery/grpc/client (interfaces: FileServiceClient)

// Package mock_files is a generated GoMock package.
package mock_files

import (
	context "context"
	io "io"
	repository_files "patreon/internal/microservices/files/files/repository/files"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileServiceClient is a mock of FileServiceClient interface.
type MockFileServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockFileServiceClientMockRecorder
}

// MockFileServiceClientMockRecorder is the mock recorder for MockFileServiceClient.
type MockFileServiceClientMockRecorder struct {
	mock *MockFileServiceClient
}

// NewMockFileServiceClient creates a new mock instance.
func NewMockFileServiceClient(ctrl *gomock.Controller) *MockFileServiceClient {
	mock := &MockFileServiceClient{ctrl: ctrl}
	mock.recorder = &MockFileServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileServiceClient) EXPECT() *MockFileServiceClientMockRecorder {
	return m.recorder
}

// LoadFile mocks base method.
func (m *MockFileServiceClient) LoadFile(arg0 context.Context, arg1 string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFile", arg0, arg1)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFile indicates an expected call of LoadFile.
func (mr *MockFileServiceClientMockRecorder) LoadFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFile", reflect.TypeOf((*MockFileServiceClient)(nil).LoadFile), arg0, arg1)
}

// SaveFile mocks base method.
func (m *MockFileServiceClient) SaveFile(arg0 context.Context, arg1 io.Reader, arg2 repository_files.FileName, arg3 repository_files.TypeFiles) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileServiceClientMockRecorder) SaveFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileServiceClient)(nil).SaveFile), arg0, arg1, arg2, arg3)
}
