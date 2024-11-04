package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// MarkPrice represents the market data for a funding rate, including the mark price and update time.
type MarkPrice struct {
	Pair  string          `json:"pair"`  // The market that the funding rate applies to, e.g., "BTC-USD.P"
	Price decimal.Decimal `json:"price"` // The current mark price as a string, e.g., "36946.34"
	Time  time.Time       `json:"time"`  // The ISO 8601 timestamp when the mark price was last updated, e.g., "2022-06-16T12:35:10.123456Z"
}
