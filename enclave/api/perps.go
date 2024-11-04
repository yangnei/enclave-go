package api

import (
	"github.com/shopspring/decimal"

	"github.com/yangnei/enclave-go/enclave/model"
)

type TransferRequest struct {
	Symbol string
	Amount decimal.Decimal
}

type GetTransferRequest struct {
	PagingAndTimeRange
}

type GetFundingRatesRequest struct {
	Market string
}

type GetFundingRateHistoryRequest struct {
	PagingAndTimeRange
	Market string
}

type GetFundingFeesRequest struct {
	PagingAndTimeRange
	Market string
}

type GetFundingFeesResponse struct {
	PageInfo    PageInfo
	FundingFees []*model.FundingFee
}

type SetStopOrderRequest struct {
	Market            string              `json:"market"`
	PositionDirection string              `json:"positionDirection"`
	Type              model.StopOrderType `json:"type"`
	TriggerPrice      decimal.Decimal     `json:"triggerPrice"`
}

type RemoveStopOrderRequest struct {
	Market string
	Type   string
}
