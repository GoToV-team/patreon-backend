// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/repository/posts (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// PostsRepository is a mock of Repository interface.
type PostsRepository struct {
	ctrl     *gomock.Controller
	recorder *PostsRepositoryMockRecorder
}

// PostsRepositoryMockRecorder is the mock recorder for PostsRepository.
type PostsRepositoryMockRecorder struct {
	mock *PostsRepository
}

// NewPostsRepository creates a new mock instance.
func NewPostsRepository(ctrl *gomock.Controller) *PostsRepository {
	mock := &PostsRepository{ctrl: ctrl}
	mock.recorder = &PostsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PostsRepository) EXPECT() *PostsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *PostsRepository) Create(arg0 *models.CreatePost) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *PostsRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*PostsRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *PostsRepository) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *PostsRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*PostsRepository)(nil).Delete), arg0)
}

// GetPost mocks base method.
func (m *PostsRepository) GetPost(arg0, arg1 int64) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *PostsRepositoryMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*PostsRepository)(nil).GetPost), arg0, arg1)
}

// GetPostCreator mocks base method.
func (m *PostsRepository) GetPostCreator(arg0 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostCreator", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostCreator indicates an expected call of GetPostCreator.
func (mr *PostsRepositoryMockRecorder) GetPostCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostCreator", reflect.TypeOf((*PostsRepository)(nil).GetPostCreator), arg0)
}

// GetPosts mocks base method.
func (m *PostsRepository) GetPosts(arg0, arg1 int64, arg2 *models.Pagination) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *PostsRepositoryMockRecorder) GetPosts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*PostsRepository)(nil).GetPosts), arg0, arg1, arg2)
}

// UpdateCoverPost mocks base method.
func (m *PostsRepository) UpdateCoverPost(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoverPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoverPost indicates an expected call of UpdateCoverPost.
func (mr *PostsRepositoryMockRecorder) UpdateCoverPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoverPost", reflect.TypeOf((*PostsRepository)(nil).UpdateCoverPost), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *PostsRepository) UpdatePost(arg0 *models.UpdatePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *PostsRepositoryMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*PostsRepository)(nil).UpdatePost), arg0)
}
