// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/adapters/rabbitmq (interfaces: ProducerRepository)
//
// Generated by this command:
//
//	mockgen --destination=producer.mock.go --package=rabbitmq . ProducerRepository
//

// Package rabbitmq is a generated GoMock package.
package rabbitmq

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	gomock "go.uber.org/mock/gomock"
)

// MockProducerRepository is a mock of ProducerRepository interface.
type MockProducerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProducerRepositoryMockRecorder
}

// MockProducerRepositoryMockRecorder is the mock recorder for MockProducerRepository.
type MockProducerRepositoryMockRecorder struct {
	mock *MockProducerRepository
}

// NewMockProducerRepository creates a new mock instance.
func NewMockProducerRepository(ctrl *gomock.Controller) *MockProducerRepository {
	mock := &MockProducerRepository{ctrl: ctrl}
	mock.recorder = &MockProducerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerRepository) EXPECT() *MockProducerRepositoryMockRecorder {
	return m.recorder
}

// ProducerDefault mocks base method.
func (m *MockProducerRepository) ProducerDefault(arg0 context.Context, arg1, arg2 string, arg3 mmodel.Queue) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducerDefault", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProducerDefault indicates an expected call of ProducerDefault.
func (mr *MockProducerRepositoryMockRecorder) ProducerDefault(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducerDefault", reflect.TypeOf((*MockProducerRepository)(nil).ProducerDefault), arg0, arg1, arg2, arg3)
}
