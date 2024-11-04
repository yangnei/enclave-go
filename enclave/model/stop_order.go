package model

import "github.com/shopspring/decimal"

type StopOrderType string

const (
	StopOrderTypeStopLoss   StopOrderType = "stopLoss"
	StopOrderTypeTakeProfit StopOrderType = "takeProfit"
)

// StopOrder represents a stop order configuration for a specific market.
type StopOrder struct {
	Market            string              `json:"market"`            // Perps market, e.g., "BTC-USD.P"
	PositionDirection PositionDirection   `json:"positionDirection"` // Position direction ("long" or "short"), e.g., "long"
	StopLoss          decimal.NullDecimal `json:"stopLoss"`          // Stop loss trigger price, if set, e.g., "8.00"
	TakeProfit        decimal.NullDecimal `json:"takeProfit"`        // Take profit trigger price, if set, or null
}
