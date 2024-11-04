package model

import "github.com/shopspring/decimal"

type Volume struct {
	Market string          `json:"market"` // Perpetual futures market (e.g., "AVAX-USD.P")
	Volume decimal.Decimal `json:"volume"` // 24-hour volume in base size (e.g., "2394.2301")
}
