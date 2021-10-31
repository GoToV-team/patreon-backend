// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/likes (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// LikesUsecase is a mock of Usecase interface.
type LikesUsecase struct {
	ctrl     *gomock.Controller
	recorder *LikesUsecaseMockRecorder
}

// LikesUsecaseMockRecorder is the mock recorder for LikesUsecase.
type LikesUsecaseMockRecorder struct {
	mock *LikesUsecase
}

// NewLikesUsecase creates a new mock instance.
func NewLikesUsecase(ctrl *gomock.Controller) *LikesUsecase {
	mock := &LikesUsecase{ctrl: ctrl}
	mock.recorder = &LikesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *LikesUsecase) EXPECT() *LikesUsecaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *LikesUsecase) Add(arg0 *models.Like) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *LikesUsecaseMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*LikesUsecase)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *LikesUsecase) Delete(arg0, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *LikesUsecaseMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*LikesUsecase)(nil).Delete), arg0, arg1)
}
