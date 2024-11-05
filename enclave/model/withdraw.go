package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WithdrawalRequest struct {
	AccountID            string          `json:"account_id"`             // Internal ID associated with the account that made the request
	Address              string          `json:"address"`                // Address to initiate withdrawal to
	Amount               decimal.Decimal `json:"amount"`                 // Amount of the target coin to withdraw
	CustomerWithdrawalID string          `json:"customer_withdrawal_id"` // Unique ID a customer associates with a withdrawal
	Symbol               string          `json:"symbol"`                 // Symbol the user wishes to withdraw
}

type WithdrawalStatus struct {
	ConfirmationNumber int64              `json:"confirmation_number,omitempty"` // Confirmation number when status is pending
	OriginalRequest    *WithdrawalRequest `json:"original_request"`              // Original withdrawal request details
	TxID               string             `json:"txid,omitempty"`                // Blockchain transaction ID
	WithdrawalID       string             `json:"withdrawal_id"`                 // ID of the withdrawal request
	WithdrawalStatus   string             `json:"withdrawal_status"`             // Status of the withdrawal
}

type Withdrawal struct {
	Address      string          `json:"address"`        // Address to which the withdrawal was initiated
	Coin         string          `json:"coin"`           // The coin which was withdrawn
	Size         decimal.Decimal `json:"size"`           // Amount of the coin withdrawn
	Status       string          `json:"status"`         // Status of the withdrawal (e.g., "WITHDRAWAL_CONFIRMED")
	Time         time.Time       `json:"time"`           // Time the withdrawal was initiated (ISO8601 format)
	TxID         string          `json:"txid,omitempty"` // Blockchain transaction ID (optional)
	WithdrawalID string          `json:"withdrawal_id"`  // System-assigned unique withdrawal ID
}

type NewWithdrawal struct {
	CustomerWithdrawalID string `json:"customer_withdrawal_id"` // User-specified ID for the withdrawal
	WithdrawalID         string `json:"withdrawal_id"`          // System-assigned ID associated with the withdrawal
	WithdrawalStatus     string `json:"withdrawal_status"`      // Status of the withdrawal
}
