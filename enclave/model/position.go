package model

import "github.com/shopspring/decimal"

type PositionDirection string

const (
	PositionDirectionLong  PositionDirection = "long"
	PositionDirectionShort PositionDirection = "short"
)

// Position represents a single position in the Perps market.
type Position struct {
	Market            string            `json:"market"`               // Perps market, e.g., "BTC-USD.P"
	Direction         PositionDirection `json:"direction"`            // Position direction, either "long" or "short"
	NetQuantity       decimal.Decimal   `json:"netQuantity"`          // Decimal with the size of the position
	AverageEntryPrice decimal.Decimal   `json:"averageEntryPrice"`    // Average entry price
	UsedMargin        decimal.Decimal   `json:"usedMargin"`           // Margin used for the position
	UnrealizedPnl     decimal.Decimal   `json:"unrealizedPnl"`        // Unrealized PNL, positive is profit, negative is loss
	MarkPrice         decimal.Decimal   `json:"markPrice"`            // Mark price
	LiquidationPrice  decimal.Decimal   `json:"liquidationPrice"`     // Price at which the position will be liquidated
	BankruptcyPrice   decimal.Decimal   `json:"bankruptcyPrice"`      // Price at which liquidation happens
	MaintenanceMargin decimal.Decimal   `json:"maintenanceMargin"`    // Maintenance margin required for this position
	StopLoss          decimal.Decimal   `json:"stopLoss,omitempty"`   // Optional: triggering price for stop loss order
	TakeProfit        decimal.Decimal   `json:"takeProfit,omitempty"` // Optional: triggering price for take profit order
}
