// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/asset (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=asset.mock.go --package=asset . Repository
//

// Package asset is a generated GoMock package.
package asset

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 *mmodel.Asset) (*mmodel.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*mmodel.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 context.Context, arg1, arg2, arg3 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Find mocks base method.
func (m *MockRepository) Find(arg0 context.Context, arg1, arg2, arg3 uuid.UUID) (*mmodel.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mmodel.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), arg0, arg1, arg2, arg3)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(arg0 context.Context, arg1, arg2 uuid.UUID, arg3, arg4 int) ([]*mmodel.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*mmodel.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), arg0, arg1, arg2, arg3, arg4)
}

// FindByNameOrCode mocks base method.
func (m *MockRepository) FindByNameOrCode(arg0 context.Context, arg1, arg2 uuid.UUID, arg3, arg4 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNameOrCode", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNameOrCode indicates an expected call of FindByNameOrCode.
func (mr *MockRepositoryMockRecorder) FindByNameOrCode(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNameOrCode", reflect.TypeOf((*MockRepository)(nil).FindByNameOrCode), arg0, arg1, arg2, arg3, arg4)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 []uuid.UUID) ([]*mmodel.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*mmodel.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockRepository) Update(arg0 context.Context, arg1, arg2, arg3 uuid.UUID, arg4 *mmodel.Asset) (*mmodel.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*mmodel.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), arg0, arg1, arg2, arg3, arg4)
}
