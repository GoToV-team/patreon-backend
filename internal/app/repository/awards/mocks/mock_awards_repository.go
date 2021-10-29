// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/repository/awards (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// AwardsRepository is a mock of Repository interface.
type AwardsRepository struct {
	ctrl     *gomock.Controller
	recorder *AwardsRepositoryMockRecorder
}

// AwardsRepositoryMockRecorder is the mock recorder for AwardsRepository.
type AwardsRepositoryMockRecorder struct {
	mock *AwardsRepository
}

// NewAwardsRepository creates a new mock instance.
func NewAwardsRepository(ctrl *gomock.Controller) *AwardsRepository {
	mock := &AwardsRepository{ctrl: ctrl}
	mock.recorder = &AwardsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *AwardsRepository) EXPECT() *AwardsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *AwardsRepository) Create(arg0 *models.Awards) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *AwardsRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*AwardsRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *AwardsRepository) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *AwardsRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*AwardsRepository)(nil).Delete), arg0)
}

// GetAwards mocks base method.
func (m *AwardsRepository) GetAwards(arg0 int64) ([]models.Awards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAwards", arg0)
	ret0, _ := ret[0].([]models.Awards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAwards indicates an expected call of GetAwards.
func (mr *AwardsRepositoryMockRecorder) GetAwards(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAwards", reflect.TypeOf((*AwardsRepository)(nil).GetAwards), arg0)
}

// GetByID mocks base method.
func (m *AwardsRepository) GetByID(arg0 int64) (*models.Awards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.Awards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *AwardsRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*AwardsRepository)(nil).GetByID), arg0)
}

// Update mocks base method.
func (m *AwardsRepository) Update(arg0 *models.Awards) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *AwardsRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*AwardsRepository)(nil).Update), arg0)
}