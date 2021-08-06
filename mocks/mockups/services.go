// Code generated by MockGen. DO NOT EDIT.
// Source: ./services.go

// Package mockups is a generated GoMock package.
package mockups

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
)

// MockCardService is a mock of CardService interface
type MockCardService struct {
	ctrl     *gomock.Controller
	recorder *MockCardServiceMockRecorder
}

// MockCardServiceMockRecorder is the mock recorder for MockCardService
type MockCardServiceMockRecorder struct {
	mock *MockCardService
}

// NewMockCardService creates a new mock instance
func NewMockCardService(ctrl *gomock.Controller) *MockCardService {
	mock := &MockCardService{ctrl: ctrl}
	mock.recorder = &MockCardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCardService) EXPECT() *MockCardServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCardService) Create(card domain.Card) (domain.Card, error) {
	ret := m.ctrl.Call(m, "Create", card)
	ret0, _ := ret[0].(domain.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCardServiceMockRecorder) Create(card interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCardService)(nil).Create), card)
}

// Get mocks base method
func (m *MockCardService) Get(id string) (domain.Card, error) {
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCardServiceMockRecorder) Get(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCardService)(nil).Get), id)
}

// MockTransactionService is a mock of TransactionService interface
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTransactionService) Create(transaction domain.Transaction) (domain.Transaction, error) {
	ret := m.ctrl.Call(m, "Create", transaction)
	ret0, _ := ret[0].(domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTransactionServiceMockRecorder) Create(transaction interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionService)(nil).Create), transaction)
}

// Get mocks base method
func (m *MockTransactionService) Get(id string) (domain.Transaction, error) {
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTransactionServiceMockRecorder) Get(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTransactionService)(nil).Get), id)
}