// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/info (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// InfoUsecase is a mock of Usecase interface.
type InfoUsecase struct {
	ctrl     *gomock.Controller
	recorder *InfoUsecaseMockRecorder
}

// InfoUsecaseMockRecorder is the mock recorder for InfoUsecase.
type InfoUsecaseMockRecorder struct {
	mock *InfoUsecase
}

// NewInfoUsecase creates a new mock instance.
func NewInfoUsecase(ctrl *gomock.Controller) *InfoUsecase {
	mock := &InfoUsecase{ctrl: ctrl}
	mock.recorder = &InfoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *InfoUsecase) EXPECT() *InfoUsecaseMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *InfoUsecase) Get() (*models.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*models.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *InfoUsecaseMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*InfoUsecase)(nil).Get))
}
