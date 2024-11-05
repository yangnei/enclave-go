package model

import "github.com/shopspring/decimal"

type WithdrawalLimit struct {
	WithdrawalLimitUsd    decimal.Decimal `json:"withdrawalLimitUsd"`    // Daily withdrawal limit in USD value
	CurrentWithdrawalsUsd decimal.Decimal `json:"currentWithdrawalsUsd"` // Amount withdrawn in the current day in USD value
}
