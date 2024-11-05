package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/shopspring/decimal"

	"github.com/yangnei/enclave-go/enclave/api"
	"github.com/yangnei/enclave-go/enclave/model"
)

type PerpsClient interface {
	GetPositions() ([]*model.Position, error)
	GetBalance() (*model.Balance, error)
	Transfer(req *api.TransferRequest) (*model.Transfer, error)
	GetTransfers(req *api.GetTransferRequest) ([]*model.Transfer, error)
	GetMarkPrices() (map[string]*model.MarkPrice, error)
	GetFundingRates(req *api.GetFundingRatesRequest) (*model.FundingRate, error)
	GetFundingRateHistory(req *api.GetFundingRateHistoryRequest) ([]*model.FundingRate, error)
	GetStopOrders() ([]*model.StopOrder, error)
	SetStopOrder(req *api.SetStopOrderRequest) ([]*model.StopOrder, error)
	RemoveStopOrder(req *api.RemoveStopOrderRequest) ([]*model.StopOrder, error)
	GetOpenInterest() ([]*model.OpenInterest, error)
	GetVolume() ([]*model.Volume, error)

	AddOrder(req *api.AddOrderRequest) (*model.Order, error)
	GetOrders(req *api.GetOrdersRequest) (*api.GetOrdersResponse, error)
	GetOrder(req *api.GetOrderRequest) (*model.Order, error)
	GetOrdersCSV(req *api.GetOrdersCSVRequest) (string, error)
	CancelOrder(req *api.CancelOrderRequest) (*model.Order, error)
	CancelOrders(req *api.CancelOrdersRequest) error
	GetDepth(req *api.GetDepthRequest) (*model.OrderBook, error)
	GetFills(req *api.GetFillsRequest) ([]*model.Fill, error)
	GetFillsByID(req *api.GetFillsByIDRequest) ([]*model.Fill, error)
	GetFillsCSV(req *api.GetFillsCSVRequest) (string, error)
}

// perpsClient provides methods specific to the perpsClient API, calling an instance of BaseClient for requests and handling authentication.
type perpsClient struct {
	BaseClient
	OrderFillClient
}

// NewPerpsClient initializes a new perpsClient client with the provided API key, API secret, and base URL.
func NewPerpsClient(apiKey, apiSecret, baseURL string) PerpsClient {
	return NewPerpsClientWithBase(NewBaseClient(apiKey, apiSecret, baseURL))
}

// NewPerpsClientWithBase initializes a new perpsClient client with the provided BaseClient.
func NewPerpsClientWithBase(baseClient BaseClient) PerpsClient {
	return &perpsClient{
		BaseClient: baseClient,
		OrderFillClient: &orderFillClient{
			BaseClient: baseClient,
			prefix:     "/v1/perps",
		},
	}
}

// GetPositions retrieves a list of all positions for all markets.
// GET /v1/perps/positions
func (p *perpsClient) GetPositions() ([]*model.Position, error) {
	resp, err := p.Get("/v1/perps/positions", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.Position]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetBalance retrieves the balance summary for the margin account.
// GET /v1/perps/balance
func (p *perpsClient) GetBalance() (*model.Balance, error) {
	resp, err := p.Get("/v1/perps/balance", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[*model.Balance]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// Transfer executes a transfer between the main wallet and margin account.
// Positive amount for deposit, negative for withdrawal.
// POST /v1/perps/transfers
func (p *perpsClient) Transfer(req *api.TransferRequest) (*model.Transfer, error) {
	transfer := &model.Transfer{
		Symbol: req.Symbol,
		Amount: req.Amount,
	}
	if req.Amount.LessThan(decimal.Zero) {
		transfer.From = &model.AccountWalletKey{
			Wallet: model.WalletMargin,
		}
		transfer.To = &model.AccountWalletKey{
			Wallet: model.WalletMain,
		}
	} else {
		transfer.From = &model.AccountWalletKey{
			Wallet: model.WalletMain,
		}
		transfer.To = &model.AccountWalletKey{
			Wallet: model.WalletMargin,
		}
	}

	requestBody, err := json.Marshal(transfer)
	if err != nil {
		return nil, err
	}

	resp, err := p.Post("/v1/perps/transfers", string(requestBody), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[*model.Transfer]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetTransfers retrieves a list of all transfers for the margin account.
// GET /v1/perps/transfers
func (p *perpsClient) GetTransfers(req *api.GetTransferRequest) ([]*model.Transfer, error) {
	query := req.GetUrlValues()
	path := "/v1/perps/transfers?" + query.Encode()

	resp, err := p.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.Transfer]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetMarkPrices retrieves the current mark price for all markets.
// GET /v1/perps/mark_prices
func (p *perpsClient) GetMarkPrices() (map[string]*model.MarkPrice, error) {
	resp, err := p.Get("/v1/perps/mark_prices", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[map[string]*model.MarkPrice]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetFundingRates retrieves the current funding rate for a given market.
// GET /v1/perps/funding_rates
func (p *perpsClient) GetFundingRates(req *api.GetFundingRatesRequest) (*model.FundingRate, error) {
	query := url.Values{}
	if req.Market == "" {
		return nil, fmt.Errorf("market is required")
	}
	query.Set("market", req.Market)

	path := "/v1/perps/funding_rates?" + query.Encode()

	resp, err := p.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[*model.FundingRate]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetFundingRateHistory retrieves the historical funding rate for a market.
// GET /v1/perps/funding_rate_history
func (p *perpsClient) GetFundingRateHistory(req *api.GetFundingRateHistoryRequest) ([]*model.FundingRate, error) {
	query := req.GetUrlValues()
	if req.Market != "" {
		query.Set("market", req.Market)
	}

	path := "/v1/perps/funding_rate_history?" + query.Encode()

	resp, err := p.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.FundingRate]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetFundingFees retrieves the historical funding fee payments in a market.
func (p *perpsClient) GetFundingFees(req *api.GetFundingFeesRequest) (*api.GetFundingFeesResponse, error) {
	query := req.GetUrlValues()
	if req.Market == "" {
		return nil, fmt.Errorf("market is required")
	}
	query.Set("market", req.Market)

	path := "/v1/perps/funding_fees?" + query.Encode()

	resp, err := p.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.PaginatedResponse[[]*model.FundingFee]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &api.GetFundingFeesResponse{
		PageInfo:    apiResp.PageInfo,
		FundingFees: apiResp.Result,
	}, nil
}

// GetStopOrders retrieves a list of all stop orders for all markets.
// GET /v1/perps/stop_order
func (p *perpsClient) GetStopOrders() ([]*model.StopOrder, error) {
	resp, err := p.Get("/v1/perps/stop_order", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.StopOrder]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// SetStopOrder creates or updates a stop order for a particular position by market and direction.
// POST /v1/perps/stop_order
func (p *perpsClient) SetStopOrder(req *api.SetStopOrderRequest) ([]*model.StopOrder, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := p.Post("/v1/perps/stop_order", string(requestBody), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.StopOrder]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// RemoveStopOrder cancels an order by either client order ID or order ID.
func (p *perpsClient) RemoveStopOrder(req *api.RemoveStopOrderRequest) ([]*model.StopOrder, error) {
	query := url.Values{}
	if req.Market == "" {
		return nil, fmt.Errorf("market is required")
	}
	query.Set("market", req.Market)

	if req.Type != "" {
		query.Set("type", req.Type)
	}

	path := "/v1/perps/stop_order?" + query.Encode()

	resp, err := p.Delete(path, "", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.StopOrder]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetOpenInterest retrieves the open interest for a given market.
// GET /v1/perps/open_interest
func (p *perpsClient) GetOpenInterest() ([]*model.OpenInterest, error) {
	resp, err := p.Get("/v1/perps/open_interest", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.OpenInterest]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}

// GetVolume retrieves the 24-hour trading volume for all markets.
// GET /v1/perps/volume
func (p *perpsClient) GetVolume() ([]*model.Volume, error) {
	resp, err := p.Get("/v1/perps/volume", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, p.HandleError(resp)
	}

	apiResp := api.Response[[]*model.Volume]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Result, nil
}
