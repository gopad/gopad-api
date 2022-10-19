// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/teams/repository/repository.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gopad/gopad-api/pkg/model"
)

// MockTeamsRepository is a mock of TeamsRepository interface.
type MockTeamsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsRepositoryMockRecorder
}

// MockTeamsRepositoryMockRecorder is the mock recorder for MockTeamsRepository.
type MockTeamsRepositoryMockRecorder struct {
	mock *MockTeamsRepository
}

// NewMockTeamsRepository creates a new mock instance.
func NewMockTeamsRepository(ctrl *gomock.Controller) *MockTeamsRepository {
	mock := &MockTeamsRepository{ctrl: ctrl}
	mock.recorder = &MockTeamsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamsRepository) EXPECT() *MockTeamsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTeamsRepository) Create(arg0 context.Context, arg1 *model.Team) (*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTeamsRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTeamsRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTeamsRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTeamsRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTeamsRepository)(nil).Delete), arg0, arg1)
}

// Exists mocks base method.
func (m *MockTeamsRepository) Exists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockTeamsRepositoryMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockTeamsRepository)(nil).Exists), arg0, arg1)
}

// List mocks base method.
func (m *MockTeamsRepository) List(arg0 context.Context) ([]*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeamsRepositoryMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeamsRepository)(nil).List), arg0)
}

// Show mocks base method.
func (m *MockTeamsRepository) Show(arg0 context.Context, arg1 string) (*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0, arg1)
	ret0, _ := ret[0].(*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockTeamsRepositoryMockRecorder) Show(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockTeamsRepository)(nil).Show), arg0, arg1)
}

// Update mocks base method.
func (m *MockTeamsRepository) Update(arg0 context.Context, arg1 *model.Team) (*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTeamsRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTeamsRepository)(nil).Update), arg0, arg1)
}