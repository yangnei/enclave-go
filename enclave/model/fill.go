package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Fill represents a fill structure.
type Fill struct {
	ClientOrderID string          `json:"clientOrderId,omitempty"`
	Fee           decimal.Decimal `json:"fee"`
	FilledCost    decimal.Decimal `json:"filledCost"`
	ID            string          `json:"id"`
	Market        string          `json:"market"`
	OrderID       string          `json:"orderId"`
	Price         decimal.Decimal `json:"price"`
	Side          OrderSide       `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Time          time.Time       `json:"time"`
}
