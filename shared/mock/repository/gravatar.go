// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iwanjunaid/basesvc/usecase/gravatar/repository (interfaces: GravatarRepository,GravatarCacheRepository)

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/iwanjunaid/basesvc/domain/model"
	reflect "reflect"
)

// MockGravatarRepository is a mock of GravatarRepository interface
type MockGravatarRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGravatarRepositoryMockRecorder
}

// MockGravatarRepositoryMockRecorder is the mock recorder for MockGravatarRepository
type MockGravatarRepositoryMockRecorder struct {
	mock *MockGravatarRepository
}

// NewMockGravatarRepository creates a new mock instance
func NewMockGravatarRepository(ctrl *gomock.Controller) *MockGravatarRepository {
	mock := &MockGravatarRepository{ctrl: ctrl}
	mock.recorder = &MockGravatarRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGravatarRepository) EXPECT() *MockGravatarRepositoryMockRecorder {
	return m.recorder
}

// AvatarURL mocks base method
func (m *MockGravatarRepository) AvatarURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvatarURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvatarURL indicates an expected call of AvatarURL
func (mr *MockGravatarRepositoryMockRecorder) AvatarURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvatarURL", reflect.TypeOf((*MockGravatarRepository)(nil).AvatarURL))
}

// GetProfile mocks base method
func (m *MockGravatarRepository) GetProfile() (*model.GravatarProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile")
	ret0, _ := ret[0].(*model.GravatarProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile
func (mr *MockGravatarRepositoryMockRecorder) GetProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockGravatarRepository)(nil).GetProfile))
}

// JSONURL mocks base method
func (m *MockGravatarRepository) JSONURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSONURL indicates an expected call of JSONURL
func (mr *MockGravatarRepositoryMockRecorder) JSONURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONURL", reflect.TypeOf((*MockGravatarRepository)(nil).JSONURL))
}

// URL mocks base method
func (m *MockGravatarRepository) URL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URL indicates an expected call of URL
func (mr *MockGravatarRepositoryMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockGravatarRepository)(nil).URL))
}

// MockGravatarCacheRepository is a mock of GravatarCacheRepository interface
type MockGravatarCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGravatarCacheRepositoryMockRecorder
}

// MockGravatarCacheRepositoryMockRecorder is the mock recorder for MockGravatarCacheRepository
type MockGravatarCacheRepositoryMockRecorder struct {
	mock *MockGravatarCacheRepository
}

// NewMockGravatarCacheRepository creates a new mock instance
func NewMockGravatarCacheRepository(ctrl *gomock.Controller) *MockGravatarCacheRepository {
	mock := &MockGravatarCacheRepository{ctrl: ctrl}
	mock.recorder = &MockGravatarCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGravatarCacheRepository) EXPECT() *MockGravatarCacheRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGravatarCacheRepository) Create(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGravatarCacheRepositoryMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGravatarCacheRepository)(nil).Create), arg0, arg1, arg2)
}

// Find mocks base method
func (m *MockGravatarCacheRepository) Find(arg0 context.Context, arg1 string) (*model.GravatarProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*model.GravatarProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockGravatarCacheRepositoryMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGravatarCacheRepository)(nil).Find), arg0, arg1)
}