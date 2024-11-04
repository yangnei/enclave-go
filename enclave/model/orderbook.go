package model

import "github.com/shopspring/decimal"

// OrderBook represents the order book structure.
type OrderBook struct {
	Asks [][]decimal.Decimal `json:"asks"`
	Bids [][]decimal.Decimal `json:"bids"`
}
