package api

import "github.com/shopspring/decimal"

type GetAssetBalanceRequest struct {
	Symbol string
}

type GetWithdrawalStatusRequest struct {
	CustomerWithdrawalId string `json:"customerWithdrawalId"`
	WithdrawalId         string `json:"withdrawalId"`
}

type GetDepositAddressesRequest struct {
	Coins []string `json:"coins"`
}

type GetDepositRequest struct {
	TxId string `json:"txId"`
}

type GetDepositsCSVRequest struct {
	TimeRange
}

type GetWithdrawalByTxIdRequest struct {
	TxId string `json:"txId"`
}

type GetWithdrawalsCSVRequest struct {
	TimeRange
}

type ProvisionAddressRequest struct {
	Symbol string `json:"symbol"`
}

type WithdrawRequest struct {
	Address              string          `json:"address"`                // The address to initiate withdrawal to
	Amount               decimal.Decimal `json:"amount"`                 // The amount of the target coin to withdraw
	CustomerWithdrawalID string          `json:"customer_withdrawal_id"` // Unique ID associated with the withdrawal
	Symbol               string          `json:"symbol"`                 // Symbol of the coin to withdraw
}
