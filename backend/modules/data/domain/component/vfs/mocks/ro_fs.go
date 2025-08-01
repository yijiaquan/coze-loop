// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/modules/data/domain/component/vfs (interfaces: ROFileSystem)
//
// Generated by this command:
//
//	mockgen -destination=mocks/ro_fs.go -package=mocks . ROFileSystem
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	fs "io/fs"
	reflect "reflect"

	vfs "github.com/coze-dev/coze-loop/backend/modules/data/domain/component/vfs"
	gomock "go.uber.org/mock/gomock"
)

// MockROFileSystem is a mock of ROFileSystem interface.
type MockROFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockROFileSystemMockRecorder
	isgomock struct{}
}

// MockROFileSystemMockRecorder is the mock recorder for MockROFileSystem.
type MockROFileSystemMockRecorder struct {
	mock *MockROFileSystem
}

// NewMockROFileSystem creates a new mock instance.
func NewMockROFileSystem(ctrl *gomock.Controller) *MockROFileSystem {
	mock := &MockROFileSystem{ctrl: ctrl}
	mock.recorder = &MockROFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockROFileSystem) EXPECT() *MockROFileSystemMockRecorder {
	return m.recorder
}

// ReadDir mocks base method.
func (m *MockROFileSystem) ReadDir(ctx context.Context, name string) ([]fs.DirEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", ctx, name)
	ret0, _ := ret[0].([]fs.DirEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockROFileSystemMockRecorder) ReadDir(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockROFileSystem)(nil).ReadDir), ctx, name)
}

// ReadFile mocks base method.
func (m *MockROFileSystem) ReadFile(ctx context.Context, name string) (vfs.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", ctx, name)
	ret0, _ := ret[0].(vfs.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockROFileSystemMockRecorder) ReadFile(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockROFileSystem)(nil).ReadFile), ctx, name)
}

// Stat mocks base method.
func (m *MockROFileSystem) Stat(ctx context.Context, name string) (fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", ctx, name)
	ret0, _ := ret[0].(fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockROFileSystemMockRecorder) Stat(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockROFileSystem)(nil).Stat), ctx, name)
}
