// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yangnei/enclave-go/enclave/client (interfaces: SpotClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/yangnei/enclave-go/enclave/api"
	model "github.com/yangnei/enclave-go/enclave/model"
)

// MockSpotClient is a mock of SpotClient interface.
type MockSpotClient struct {
	ctrl     *gomock.Controller
	recorder *MockSpotClientMockRecorder
}

// MockSpotClientMockRecorder is the mock recorder for MockSpotClient.
type MockSpotClientMockRecorder struct {
	mock *MockSpotClient
}

// NewMockSpotClient creates a new mock instance.
func NewMockSpotClient(ctrl *gomock.Controller) *MockSpotClient {
	mock := &MockSpotClient{ctrl: ctrl}
	mock.recorder = &MockSpotClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpotClient) EXPECT() *MockSpotClientMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockSpotClient) AddOrder(arg0 *api.AddOrderRequest) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", arg0)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockSpotClientMockRecorder) AddOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockSpotClient)(nil).AddOrder), arg0)
}

// CancelOrder mocks base method.
func (m *MockSpotClient) CancelOrder(arg0 *api.CancelOrderRequest) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockSpotClientMockRecorder) CancelOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockSpotClient)(nil).CancelOrder), arg0)
}

// CancelOrders mocks base method.
func (m *MockSpotClient) CancelOrders(arg0 *api.CancelOrdersRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrders", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrders indicates an expected call of CancelOrders.
func (mr *MockSpotClientMockRecorder) CancelOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrders", reflect.TypeOf((*MockSpotClient)(nil).CancelOrders), arg0)
}

// GetDepth mocks base method.
func (m *MockSpotClient) GetDepth(arg0 *api.GetDepthRequest) (*model.OrderBook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDepth", arg0)
	ret0, _ := ret[0].(*model.OrderBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDepth indicates an expected call of GetDepth.
func (mr *MockSpotClientMockRecorder) GetDepth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDepth", reflect.TypeOf((*MockSpotClient)(nil).GetDepth), arg0)
}

// GetFills mocks base method.
func (m *MockSpotClient) GetFills(arg0 *api.GetFillsRequest) ([]*model.Fill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFills", arg0)
	ret0, _ := ret[0].([]*model.Fill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFills indicates an expected call of GetFills.
func (mr *MockSpotClientMockRecorder) GetFills(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFills", reflect.TypeOf((*MockSpotClient)(nil).GetFills), arg0)
}

// GetFillsByID mocks base method.
func (m *MockSpotClient) GetFillsByID(arg0 *api.GetFillsByIDRequest) ([]*model.Fill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFillsByID", arg0)
	ret0, _ := ret[0].([]*model.Fill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFillsByID indicates an expected call of GetFillsByID.
func (mr *MockSpotClientMockRecorder) GetFillsByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFillsByID", reflect.TypeOf((*MockSpotClient)(nil).GetFillsByID), arg0)
}

// GetFillsCSV mocks base method.
func (m *MockSpotClient) GetFillsCSV(arg0 *api.GetFillsCSVRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFillsCSV", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFillsCSV indicates an expected call of GetFillsCSV.
func (mr *MockSpotClientMockRecorder) GetFillsCSV(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFillsCSV", reflect.TypeOf((*MockSpotClient)(nil).GetFillsCSV), arg0)
}

// GetOrder mocks base method.
func (m *MockSpotClient) GetOrder(arg0 *api.GetOrderRequest) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockSpotClientMockRecorder) GetOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockSpotClient)(nil).GetOrder), arg0)
}

// GetOrders mocks base method.
func (m *MockSpotClient) GetOrders(arg0 *api.GetOrdersRequest) (*api.GetOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].(*api.GetOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockSpotClientMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockSpotClient)(nil).GetOrders), arg0)
}

// GetOrdersCSV mocks base method.
func (m *MockSpotClient) GetOrdersCSV(arg0 *api.GetOrdersCSVRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersCSV", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersCSV indicates an expected call of GetOrdersCSV.
func (mr *MockSpotClientMockRecorder) GetOrdersCSV(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersCSV", reflect.TypeOf((*MockSpotClient)(nil).GetOrdersCSV), arg0)
}
