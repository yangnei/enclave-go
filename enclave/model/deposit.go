package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Deposit struct {
	Coin                  string          `json:"coin"`                  // The coin which was deposited
	CurrentConfirmations  int64           `json:"currentConfirmations"`  // Current number of confirmations if the deposit is pending
	RequiredConfirmations int64           `json:"requiredConfirmations"` // Required number of confirmations for confirmation
	Size                  decimal.Decimal `json:"size"`                  // Amount of the coin deposited
	Status                string          `json:"status"`                // Status of the deposit
	Time                  time.Time       `json:"time"`                  // Time the deposit was made (ISO8601 format)
	TxID                  string          `json:"txid"`                  // Transaction ID of the deposit
}
