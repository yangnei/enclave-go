package model

import "github.com/shopspring/decimal"

type OpenInterest struct {
	Market        string          `json:"market"`        // Perpetual futures market (e.g., "AVAX-USD.P")
	OpenInterest  decimal.Decimal `json:"openInterest"`  // Open interest in base size (e.g., "2394.2301")
	NotionalValue decimal.Decimal `json:"notionalValue"` // Notional value of the open interest at the current mark price (e.g., "24311.72")
}
