// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/repository/creator (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// CreatorRepository is a mock of Repository interface.
type CreatorRepository struct {
	ctrl     *gomock.Controller
	recorder *CreatorRepositoryMockRecorder
}

// CreatorRepositoryMockRecorder is the mock recorder for CreatorRepository.
type CreatorRepositoryMockRecorder struct {
	mock *CreatorRepository
}

// NewCreatorRepository creates a new mock instance.
func NewCreatorRepository(ctrl *gomock.Controller) *CreatorRepository {
	mock := &CreatorRepository{ctrl: ctrl}
	mock.recorder = &CreatorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CreatorRepository) EXPECT() *CreatorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *CreatorRepository) Create(arg0 *models.Creator) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *CreatorRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CreatorRepository)(nil).Create), arg0)
}

// ExistsCreator mocks base method.
func (m *CreatorRepository) ExistsCreator(arg0 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsCreator", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsCreator indicates an expected call of ExistsCreator.
func (mr *CreatorRepositoryMockRecorder) ExistsCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsCreator", reflect.TypeOf((*CreatorRepository)(nil).ExistsCreator), arg0)
}

// GetCreator mocks base method.
func (m *CreatorRepository) GetCreator(arg0 int64) (*models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator", arg0)
	ret0, _ := ret[0].(*models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *CreatorRepositoryMockRecorder) GetCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*CreatorRepository)(nil).GetCreator), arg0)
}

// GetCreators mocks base method.
func (m *CreatorRepository) GetCreators() ([]models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreators")
	ret0, _ := ret[0].([]models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreators indicates an expected call of GetCreators.
func (mr *CreatorRepositoryMockRecorder) GetCreators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreators", reflect.TypeOf((*CreatorRepository)(nil).GetCreators))
}

// UpdateAvatar mocks base method.
func (m *CreatorRepository) UpdateAvatar(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *CreatorRepositoryMockRecorder) UpdateAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*CreatorRepository)(nil).UpdateAvatar), arg0, arg1)
}

// UpdateCover mocks base method.
func (m *CreatorRepository) UpdateCover(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCover", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCover indicates an expected call of UpdateCover.
func (mr *CreatorRepositoryMockRecorder) UpdateCover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCover", reflect.TypeOf((*CreatorRepository)(nil).UpdateCover), arg0, arg1)
}
