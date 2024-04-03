// Code generated by MockGen. DO NOT EDIT.
// Source: item.go
//
// Generated by this command:
//
//	mockgen -destination=item_mock.go -source=item.go -package=server
//

// Package server is a generated GoMock package.
package server

import (
	reflect "reflect"

	model "github.com/0x2e/fusion/model"
	repo "github.com/0x2e/fusion/repo"
	gomock "go.uber.org/mock/gomock"
)

// MockItemRepo is a mock of ItemRepo interface.
type MockItemRepo struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepoMockRecorder
}

// MockItemRepoMockRecorder is the mock recorder for MockItemRepo.
type MockItemRepoMockRecorder struct {
	mock *MockItemRepo
}

// NewMockItemRepo creates a new mock instance.
func NewMockItemRepo(ctrl *gomock.Controller) *MockItemRepo {
	mock := &MockItemRepo{ctrl: ctrl}
	mock.recorder = &MockItemRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepo) EXPECT() *MockItemRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockItemRepo) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockItemRepoMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemRepo)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockItemRepo) Get(id uint) (*model.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockItemRepoMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockItemRepo)(nil).Get), id)
}

// List mocks base method.
func (m *MockItemRepo) List(filter repo.ItemFilter, page, pageSize int) ([]*model.Item, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filter, page, pageSize)
	ret0, _ := ret[0].([]*model.Item)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockItemRepoMockRecorder) List(filter, page, pageSize any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockItemRepo)(nil).List), filter, page, pageSize)
}

// UpdateBookmark mocks base method.
func (m *MockItemRepo) UpdateBookmark(id uint, bookmark *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookmark", id, bookmark)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBookmark indicates an expected call of UpdateBookmark.
func (mr *MockItemRepoMockRecorder) UpdateBookmark(id, bookmark any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookmark", reflect.TypeOf((*MockItemRepo)(nil).UpdateBookmark), id, bookmark)
}

// UpdateUnread mocks base method.
func (m *MockItemRepo) UpdateUnread(ids []uint, unread *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnread", ids, unread)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnread indicates an expected call of UpdateUnread.
func (mr *MockItemRepoMockRecorder) UpdateUnread(ids, unread any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnread", reflect.TypeOf((*MockItemRepo)(nil).UpdateUnread), ids, unread)
}
