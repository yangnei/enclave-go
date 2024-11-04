package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransferType int
type Wallet string

const (
	TransferTypeMargin     TransferType = 0
	TransferTypeSubAccount TransferType = 1

	WalletMain   Wallet = "main"
	WalletMargin Wallet = "margin"
)

// AccountWalletKey represents the wallet and account information involved in the transfer.
type AccountWalletKey struct {
	ID     string `json:"id"`     // ID of the account/wallet
	Wallet Wallet `json:"wallet"` // Wallet type, e.g., "main" or "margin"
}

// Transfer represents a transfer record with details about the transfer.
type Transfer struct {
	ID     string            `json:"id"`     // ID of the transfer
	From   *AccountWalletKey `json:"from"`   // The account/wallet the transfer pulled funds from
	To     *AccountWalletKey `json:"to"`     // The account/wallet the transfer sent funds to
	Amount decimal.Decimal   `json:"amount"` // The amount transferred
	Symbol string            `json:"symbol"` // The symbol transferred, e.g., "USDC"
	Time   time.Time         `json:"time"`   // ISO8601 timestamp of the transfer creation time
	Type   TransferType      `json:"type"`   // Transfer type: 0 for margin, 1 for subaccount
}
