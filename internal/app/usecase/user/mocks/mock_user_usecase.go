// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/user (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	io "io"
	models "patreon/internal/app/models"
	repository_files "patreon/internal/app/repository/files"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// UserUsecase is a mock of Usecase interface.
type UserUsecase struct {
	ctrl     *gomock.Controller
	recorder *UserUsecaseMockRecorder
}

// UserUsecaseMockRecorder is the mock recorder for UserUsecase.
type UserUsecaseMockRecorder struct {
	mock *UserUsecase
}

// NewUserUsecase creates a new mock instance.
func NewUserUsecase(ctrl *gomock.Controller) *UserUsecase {
	mock := &UserUsecase{ctrl: ctrl}
	mock.recorder = &UserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserUsecase) EXPECT() *UserUsecaseMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *UserUsecase) Check(arg0, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *UserUsecaseMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*UserUsecase)(nil).Check), arg0, arg1)
}

// Create mocks base method.
func (m *UserUsecase) Create(arg0 *models.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *UserUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*UserUsecase)(nil).Create), arg0)
}

// GetProfile mocks base method.
func (m *UserUsecase) GetProfile(arg0 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *UserUsecaseMockRecorder) GetProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*UserUsecase)(nil).GetProfile), arg0)
}

// UpdateAvatar mocks base method.
func (m *UserUsecase) UpdateAvatar(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *UserUsecaseMockRecorder) UpdateAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*UserUsecase)(nil).UpdateAvatar), arg0, arg1, arg2)
}

// UpdatePassword mocks base method.
func (m *UserUsecase) UpdatePassword(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *UserUsecaseMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*UserUsecase)(nil).UpdatePassword), arg0, arg1)
}
