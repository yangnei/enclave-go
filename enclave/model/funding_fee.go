package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// FundingFee represents details about a funding fee paid in a specific market.
type FundingFee struct {
	Market            string            `json:"market"`            // The market where the fee was paid, e.g., "AVAX-USD.P"
	Rate              decimal.Decimal   `json:"rate"`              // The funding rate that led to this fee, e.g., "0.0000125"
	Time              time.Time         `json:"time"`              // ISO 8601 timestamp of when the fee was paid, e.g., "2022-06-16T12:35:10.123456Z"
	Amount            decimal.Decimal   `json:"amount"`            // Actual amount of USDC transferred, positive if earned, negative if paid, e.g., "-0.16402"
	Payer             PositionDirection `json:"payer"`             // The side that paid the fee, either "long" or "short"
	MarkPrice         decimal.Decimal   `json:"markPrice"`         // The mark price at the time the fee was paid, e.g., "12.03"
	PositionSize      decimal.Decimal   `json:"positionSize"`      // The base size of the position at the time the fee was paid, e.g., "1.00"
	PositionDirection PositionDirection `json:"positionDirection"` // The direction of the position at the time the fee was paid, e.g., "long"
}
