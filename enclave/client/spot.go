package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/yangnei/enclave-go/enclave/api"
	"github.com/yangnei/enclave-go/enclave/model"
)

type SpotClient interface {
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

// spotClient contains the spotClient-specific API and calls an instance of BaseClient to make requests and handle auth.
type spotClient struct {
	BaseClient
}

// NewSpotClientWithBase initializes a new spotClient client with the provided baseClient.
func NewSpotClientWithBase(baseClient BaseClient) SpotClient {
	return &spotClient{
		BaseClient: baseClient,
	}
}

// NewSpotClient initializes a new spotClient client with the provided API key and secret.
func NewSpotClient(apiKey, apiSecret, baseURL string) SpotClient {
	return NewSpotClientWithBase(NewBaseClient(apiKey, apiSecret, baseURL))
}

// AddOrder creates a spot order.
// POST /v1/orders
func (s *spotClient) AddOrder(req *api.AddOrderRequest) (*model.Order, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.Post("/v1/orders", string(requestBody), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[*model.Order]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// GetOrders retrieves orders that meet the optional parameters.
// GET /v1/orders
func (s *spotClient) GetOrders(req *api.GetOrdersRequest) (*api.GetOrdersResponse, error) {
	query := req.GetUrlValues()
	if req.Market != "" {
		query.Set("market", req.Market)
	}
	if req.Status != "" {
		query.Set("status", string(req.Status))
	}

	path := "/v1/orders"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.PaginatedResponse[[]model.Order]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return &api.GetOrdersResponse{
		PageInfo: apiResp.PageInfo,
		Orders:   apiResp.Result,
	}, nil
}

// GetOrder retrieves an order by client order ID or internal order ID.
// Exactly one of clientOrderID or orderID must be provided.
// GET /v1/orders/{orderID} or GET /v1/orders/client:{clientOrderID}
func (s *spotClient) GetOrder(req *api.GetOrderRequest) (*model.Order, error) {
	if (req.ClientOrderID == "" && req.OrderID == "") || (req.ClientOrderID != "" && req.OrderID != "") {
		return nil, fmt.Errorf("must provide exactly one of clientOrderID or orderID")
	}

	var path string
	if req.ClientOrderID != "" {
		path = fmt.Sprintf("/v1/orders/client:%s", req.ClientOrderID)
	} else {
		path = fmt.Sprintf("/v1/orders/%s", req.OrderID)
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[*model.Order]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// GetOrdersCSV retrieves orders in CSV format that meet the optional parameters.
// GET /v1/orders/csv
func (s *spotClient) GetOrdersCSV(req *api.GetOrdersCSVRequest) (string, error) {
	query := req.GetUrlValues()
	if req.Market != "" {
		query.Set("market", req.Market)
	}
	if req.Status != "" {
		query.Set("status", string(req.Status))
	}

	path := "/v1/orders/csv"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", s.HandleError(resp)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// CancelOrder cancels an order by client order ID or internal order ID.
// Exactly one of clientOrderID or orderID must be provided.
// DELETE /v1/orders/{orderID} or DELETE /v1/orders/client:{clientOrderID}
func (s *spotClient) CancelOrder(req *api.CancelOrderRequest) (*model.Order, error) {
	if (req.ClientOrderID == "" && req.OrderID == "") || (req.ClientOrderID != "" && req.OrderID != "") {
		return nil, fmt.Errorf("must provide exactly one of clientOrderID or orderID")
	}

	var path string
	if req.ClientOrderID != "" {
		path = fmt.Sprintf("/v1/orders/client:%s", req.ClientOrderID)
	} else {
		path = fmt.Sprintf("/v1/orders/%s", req.OrderID)
	}

	resp, err := s.Delete(path, "", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[*model.Order]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// CancelOrders cancels all orders, optionally per market.
// DELETE /v1/orders
func (s *spotClient) CancelOrders(req *api.CancelOrdersRequest) error {
	query := url.Values{}
	if req.Market != "" {
		query.Set("market", req.Market)
	}

	path := "/v1/orders"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Delete(path, "", nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return s.HandleError(resp)
	}

	apiResp := api.Response[struct{}]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}
	if !apiResp.Success {
		return fmt.Errorf("API error: %s", apiResp.Error)
	}

	return nil
}

// GetDepth returns the order book in a market, optionally to a specified depth.
// GET /v1/depth
func (s *spotClient) GetDepth(req *api.GetDepthRequest) (*model.OrderBook, error) {
	query := url.Values{}
	if req.Market == "" {
		return nil, fmt.Errorf("market is required")
	}
	query.Set("market", req.Market)

	if req.Depth != 0 {
		query.Set("depth", strconv.Itoa(req.Depth))
	}

	path := "/v1/depth"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[*model.OrderBook]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// GetFills retrieves fills that meet the optional parameters.
// GET /v1/fills
func (s *spotClient) GetFills(req *api.GetFillsRequest) ([]*model.Fill, error) {
	query := req.GetUrlValues()
	if req.Market != "" {
		query.Set("market", req.Market)
	}

	path := "/v1/fills"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[[]*model.Fill]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// GetFillsByID retrieves fills by client order ID or internal order ID.
// Exactly one of clientOrderID or orderID must be provided.
// GET /v1/fills/client:{clientOrderID} or GET /v1/orders/{orderID}/fills
func (s *spotClient) GetFillsByID(req *api.GetFillsByIDRequest) ([]*model.Fill, error) {
	if (req.ClientOrderID == "" && req.OrderID == "") || (req.ClientOrderID != "" && req.OrderID != "") {
		return nil, fmt.Errorf("must provide exactly one of clientOrderID or orderID")
	}

	var path string
	if req.ClientOrderID != "" {
		path = fmt.Sprintf("/v1/fills/client:%s", req.ClientOrderID)
	} else {
		path = fmt.Sprintf("/v1/orders/%s/fills", req.OrderID)
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, s.HandleError(resp)
	}

	apiResp := api.Response[[]*model.Fill]{}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Result, nil
}

// GetFillsCSV retrieves fills in CSV format that meet the optional parameters.
// GET /v1/fills/csv
func (s *spotClient) GetFillsCSV(req *api.GetFillsCSVRequest) (string, error) {
	query := req.GetUrlValues()
	if req.Market != "" {
		query.Set("market", req.Market)
	}

	path := "/v1/fills/csv"
	if encodedQuery := query.Encode(); encodedQuery != "" {
		path += "?" + encodedQuery
	}

	resp, err := s.Get(path, nil, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", s.HandleError(resp)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
