// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/resto/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRestoUsecase is a mock of Usecase interface.
type MockRestoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRestoUsecaseMockRecorder
}

// MockRestoUsecaseMockRecorder is the mock recorder for MockRestoUsecase.
type MockRestoUsecaseMockRecorder struct {
	mock *MockRestoUsecase
}

// NewMockRestoUsecase creates a new mock instance.
func NewMockRestoUsecase(ctrl *gomock.Controller) *MockRestoUsecase {
	mock := &MockRestoUsecase{ctrl: ctrl}
	mock.recorder = &MockRestoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoUsecase) EXPECT() *MockRestoUsecaseMockRecorder {
	return m.recorder
}

// AddNewMenu mocks base method.
func (m *MockRestoUsecase) AddNewMenu(ctx context.Context, menu model.MenuItem) (model.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewMenu", ctx, menu)
	ret0, _ := ret[0].(model.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewMenu indicates an expected call of AddNewMenu.
func (mr *MockRestoUsecaseMockRecorder) AddNewMenu(ctx, menu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewMenu", reflect.TypeOf((*MockRestoUsecase)(nil).AddNewMenu), ctx, menu)
}

// CheckSession mocks base method.
func (m *MockRestoUsecase) CheckSession(ctx context.Context, data model.UserSession) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockRestoUsecaseMockRecorder) CheckSession(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockRestoUsecase)(nil).CheckSession), ctx, data)
}

// GetAllOrdersInfo mocks base method.
func (m *MockRestoUsecase) GetAllOrdersInfo(ctx context.Context) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrdersInfo", ctx)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrdersInfo indicates an expected call of GetAllOrdersInfo.
func (mr *MockRestoUsecaseMockRecorder) GetAllOrdersInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrdersInfo", reflect.TypeOf((*MockRestoUsecase)(nil).GetAllOrdersInfo), ctx)
}

// GetMenuList mocks base method.
func (m *MockRestoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuList", ctx, menuType)
	ret0, _ := ret[0].([]model.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuList indicates an expected call of GetMenuList.
func (mr *MockRestoUsecaseMockRecorder) GetMenuList(ctx, menuType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuList", reflect.TypeOf((*MockRestoUsecase)(nil).GetMenuList), ctx, menuType)
}

// GetOrderInfo mocks base method.
func (m *MockRestoUsecase) GetOrderInfo(ctx context.Context, req model.GetOrderInfoRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderInfo", ctx, req)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderInfo indicates an expected call of GetOrderInfo.
func (mr *MockRestoUsecaseMockRecorder) GetOrderInfo(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderInfo", reflect.TypeOf((*MockRestoUsecase)(nil).GetOrderInfo), ctx, req)
}

// Login mocks base method.
func (m *MockRestoUsecase) Login(ctx context.Context, req model.LoginRequest) (model.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req)
	ret0, _ := ret[0].(model.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockRestoUsecaseMockRecorder) Login(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRestoUsecase)(nil).Login), ctx, req)
}

// Order mocks base method.
func (m *MockRestoUsecase) Order(ctx context.Context, req model.OrderMenuRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", ctx, req)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Order indicates an expected call of Order.
func (mr *MockRestoUsecaseMockRecorder) Order(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockRestoUsecase)(nil).Order), ctx, req)
}

// RegisterUser mocks base method.
func (m *MockRestoUsecase) RegisterUser(ctx context.Context, req model.RegisterRequest) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, req)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRestoUsecaseMockRecorder) RegisterUser(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRestoUsecase)(nil).RegisterUser), ctx, req)
}