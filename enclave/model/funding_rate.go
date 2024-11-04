package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// FundingRate represents the estimated funding rate data for a specific market.
type FundingRate struct {
	Market       string          `json:"market"`       // The market that the funding rate applies to, e.g., "BTC-USD.P"
	Rate         decimal.Decimal `json:"rate"`         // The estimated funding rate, e.g., "0.0000125"
	IntervalEnds time.Time       `json:"intervalEnds"` // The ISO 8601 timestamp when the current funding interval will end, e.g., "2022-06-16T12:35:10.123456Z"
}
