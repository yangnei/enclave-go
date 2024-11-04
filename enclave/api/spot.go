package api

import (
	"github.com/shopspring/decimal"

	"github.com/yangnei/enclave-go/enclave/model"
)

// AddOrderRequest represents the request body for adding an order.
type AddOrderRequest struct {
	ClientOrderID string
	Market        string
	Price         decimal.Decimal
	QuoteSize     decimal.Decimal
	Side          model.OrderSide
	Size          decimal.Decimal
	Type          model.OrderType
	TimeInForce   model.TimeInForce
	PostOnly      bool
}

// GetOrdersRequest holds optional parameters for GetOrders method.
type GetOrdersRequest struct {
	PagingAndTimeRange
	Market string
	Status model.OrderStatus
}

// GetOrdersResponse represents the response from GetOrders.
type GetOrdersResponse struct {
	PageInfo PageInfo
	Orders   []model.Order
}

// GetOrderRequest holds optional parameters for GetOrder method.
type GetOrderRequest struct {
	ClientOrderID string
	OrderID       string
}

// GetOrdersCSVRequest holds optional parameters for GetOrdersCSV method.
type GetOrdersCSVRequest struct {
	TimeRange
	Market string
	Status model.OrderStatus
}

// CancelOrderRequest represents the request body for canceling an order.
type CancelOrderRequest struct {
	ClientOrderID string
	OrderID       string
}

// CancelOrdersRequest represents the request body for canceling orders.
type CancelOrdersRequest struct {
	Market string
}

// GetDepthRequest holds optional parameters for GetDepth method.
type GetDepthRequest struct {
	Market string
	Depth  int
}

// GetFillsRequest holds optional parameters for GetFills method.
type GetFillsRequest struct {
	PagingAndTimeRange
	Market string
}

// GetFillsByIDRequest represents the request body for getting fills by ID.
type GetFillsByIDRequest struct {
	ClientOrderID string
	OrderID       string
}

// GetFillsCSVRequest holds optional parameters for GetFillsCSV method.
type GetFillsCSVRequest struct {
	TimeRange
	Market string
}
