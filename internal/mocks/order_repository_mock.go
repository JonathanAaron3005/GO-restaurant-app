// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of Repository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, order)
}

// GetAllOrders mocks base method.
func (m *MockOrderRepository) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders", ctx)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockOrderRepositoryMockRecorder) GetAllOrders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockOrderRepository)(nil).GetAllOrders), ctx)
}

// GetOrderInfo mocks base method.
func (m *MockOrderRepository) GetOrderInfo(ctx context.Context, orderID string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderInfo", ctx, orderID)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderInfo indicates an expected call of GetOrderInfo.
func (mr *MockOrderRepositoryMockRecorder) GetOrderInfo(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderInfo", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderInfo), ctx, orderID)
}
