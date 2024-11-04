package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderSide string
type OrderType string
type OrderStatus string
type TimeInForce string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"

	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"

	OrderStatusOpen        OrderStatus = "open"
	OrderStatusFullyFilled OrderStatus = "fullyfilled"
	OrderStatusCanceled    OrderStatus = "canceled"

	TimeInForceGTC TimeInForce = "GTC"
	TimeInForceIOC TimeInForce = "IOC"
)

// Order represents an order structure.
type Order struct {
	CanceledAt    time.Time       `json:"canceledAt,omitempty"`
	ClientOrderID string          `json:"clientOrderId,omitempty"`
	CreatedAt     time.Time       `json:"createdAt"`
	Fee           decimal.Decimal `json:"fee"`
	FilledAt      time.Time       `json:"filledAt,omitempty"`
	FilledCost    decimal.Decimal `json:"filledCost"`
	FilledSize    decimal.Decimal `json:"filledSize"`
	Market        string          `json:"market"`
	OrderID       string          `json:"orderId"`
	Price         decimal.Decimal `json:"price"`
	Side          OrderSide       `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Status        OrderStatus     `json:"status"`
	Type          OrderType       `json:"type,omitempty"`
	TimeInForce   TimeInForce     `json:"timeInForce,omitempty"`
	CancelReason  string          `json:"cancelReason,omitempty"`
}
