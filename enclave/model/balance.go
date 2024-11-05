package model

import "github.com/shopspring/decimal"

// Balance represents the balance details of a Perps account.
type Balance struct {
	WalletBalance      decimal.Decimal `json:"walletBalance"`      // Wallet balance in USDC
	WithdrawableMargin decimal.Decimal `json:"withdrawableMargin"` // Withdrawable margin in USDC
	RealizedPnl        decimal.Decimal `json:"realizedPnl"`        // Realized PNL in USDC
	UnrealizedPnl      decimal.Decimal `json:"unrealizedPnl"`      // Unrealized PNL in USDC at mark price for open positions
	UsedMargin         decimal.Decimal `json:"usedMargin"`         // Used margin in USDC
	AvailableMargin    decimal.Decimal `json:"availableMargin"`    // Available margin in USDC
	MarginBalance      decimal.Decimal `json:"marginBalance"`      // Margin balance in USDC
	MarginRatio        decimal.Decimal `json:"marginRatio"`        // Margin ratio as a percentage
	Leverage           decimal.Decimal `json:"leverage"`           // Effective leverage for the account
	UnderLiquidation   bool            `json:"underLiquidation"`   // Whether the account is under liquidation
}

type AssetBalance struct {
	AccountID       string `json:"accountId"`       // Account ID of the user
	FreeBalance     string `json:"freeBalance"`     // Free balance of the coin
	ReservedBalance string `json:"reservedBalance"` // Reserved balance held in open orders
	Symbol          string `json:"symbol"`          // Symbol of the coin
	TotalBalance    string `json:"totalBalance"`    // Total balance of the coin
}
