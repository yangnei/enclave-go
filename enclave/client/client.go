package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yangnei/enclave-go/enclave/api"
	"github.com/yangnei/enclave-go/enclave/model"
)

type Client interface {
	SpotClient() SpotClient
	PerpsClient() PerpsClient

	Hello() (*model.Hello, error)
	AuthenticatedHello() (*model.AuthenticatedHello, error)
	GetAccount() (*model.Account, error)
	GetAddressBook() (*model.AddressBook, error)
	GetMarkets() (*model.Market, error)
	GetAssetBalance(req *api.GetAssetBalanceRequest) (*model.AssetBalance, error)
	GetWithdrawalStatus(req *api.GetWithdrawalStatusRequest) (*model.WithdrawalStatus, error)
	GetAssetBalances() ([]*model.AssetBalance, error)
	GetDepositAddresses(req *api.GetDepositAddressesRequest) ([]*model.Address, error)
	GetDeposits() ([]*model.Deposit, error)
	GetDeposit(req *api.GetDepositRequest) (*model.Deposit, error)
	GetDepositsCSV(req *api.GetDepositsCSVRequest) (string, error)
	GetWithdrawals() ([]*model.Withdrawal, error)
	GetWithdrawal() (*model.Withdrawal, error)
	GetWithdrawalLimit() (*model.WithdrawalLimit, error)
	GetWithdrawalByTxId(req *api.GetWithdrawalByTxIdRequest) (*model.Withdrawal, error)
	GetWithdrawalsCSV(req *api.GetWithdrawalsCSVRequest) (string, error)
	ProvisionAddress(req *api.ProvisionAddressRequest) (*model.Address, error)
	Withdraw(req *api.WithdrawRequest) (*model.NewWithdrawal, error)
}

type client struct {
	BaseClient
	sc SpotClient
	pc PerpsClient
}

func NewClient(apiKey, apiSecret, baseURL string) Client {
	return NewClientWithBase(NewBaseClient(apiKey, apiSecret, baseURL))
}

func NewClientWithBase(base BaseClient) Client {
	return &client{
		BaseClient: base,
		sc:         NewSpotClientWithBase(base),
		pc:         NewPerpsClientWithBase(base),
	}
}

// SpotClient returns the spot client.
func (c *client) SpotClient() SpotClient {
	return c.sc
}

// PerpsClient returns the perps client.
func (c *client) PerpsClient() PerpsClient {
	return c.pc
}

// Hello returns the server's greeting message.
func (c *client) Hello() (*model.Hello, error) {
	resp, err := c.Get("/hello", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get hello: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.HandleError(resp)
	}

	var hello model.Hello
	if err := json.NewDecoder(resp.Body).Decode(&hello); err != nil {
		return nil, fmt.Errorf("failed to decode hello: %w", err)
	}

	return &hello, nil
}

// AuthenticatedHello returns the server's greeting message for authenticated clients.
func (c *client) AuthenticatedHello() (*model.AuthenticatedHello, error) {
	resp, err := c.Get("/authedHello", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated hello: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.HandleError(resp)
	}

	var hello model.AuthenticatedHello
	if err := json.NewDecoder(resp.Body).Decode(&hello); err != nil {
		return nil, fmt.Errorf("failed to decode authenticated hello: %w", err)
	}

	return &hello, nil
}

// GetAccount returns the account associated with the API key.
func (c *client) GetAccount() (*model.Account, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetAddressBook returns the user's address book.
func (c *client) GetAddressBook() (*model.AddressBook, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetMarkets returns the list of markets available on the exchange.
func (c *client) GetMarkets() (*model.Market, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetAssetBalance returns the balance of a specific asset.
func (c *client) GetAssetBalance(req *api.GetAssetBalanceRequest) (*model.AssetBalance, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetWithdrawalStatus returns the status of a withdrawal.
func (c *client) GetWithdrawalStatus(req *api.GetWithdrawalStatusRequest) (*model.WithdrawalStatus, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetAssetBalances returns the balances of all assets.
func (c *client) GetAssetBalances() ([]*model.AssetBalance, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetDepositAddresses returns the deposit addresses for the specified coins.
func (c *client) GetDepositAddresses(req *api.GetDepositAddressesRequest) ([]*model.Address, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetDeposits returns the list of deposits.
func (c *client) GetDeposits() ([]*model.Deposit, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetDeposit returns the details of a specific deposit.
func (c *client) GetDeposit(req *api.GetDepositRequest) (*model.Deposit, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetDepositsCSV returns the list of deposits in CSV format.
func (c *client) GetDepositsCSV(req *api.GetDepositsCSVRequest) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// GetWithdrawals returns the list of withdrawals.
func (c *client) GetWithdrawals() ([]*model.Withdrawal, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetWithdrawal returns the details of a specific withdrawal.
func (c *client) GetWithdrawal() (*model.Withdrawal, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetWithdrawalLimit returns the withdrawal limits for the account.
func (c *client) GetWithdrawalLimit() (*model.WithdrawalLimit, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetWithdrawalByTxId returns the details of a withdrawal by transaction ID.
func (c *client) GetWithdrawalByTxId(req *api.GetWithdrawalByTxIdRequest) (*model.Withdrawal, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetWithdrawalsCSV returns the list of withdrawals in CSV format.
func (c *client) GetWithdrawalsCSV(req *api.GetWithdrawalsCSVRequest) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// ProvisionAddress provisions a new deposit address for the specified coin.
func (c *client) ProvisionAddress(req *api.ProvisionAddressRequest) (*model.Address, error) {
	return nil, fmt.Errorf("not implemented")
}

// Withdraw initiates a withdrawal to the specified address.
func (c *client) Withdraw(req *api.WithdrawRequest) (*model.NewWithdrawal, error) {
	return nil, fmt.Errorf("not implemented")
}
