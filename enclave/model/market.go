package model

import "github.com/shopspring/decimal"

type CurrencyPair struct {
	Base  string `json:"base"`  // Base coin of the currency pair (e.g., "AVAX")
	Quote string `json:"quote"` // Quote coin of the currency pair (e.g., "ETH")
}

type V1CrossMarketsResult struct {
	DecimalPlaces int           `json:"decimalPlaces"`      // Number of decimal places for the pair
	Pair          *CurrencyPair `json:"pair"`               // Currency pair information
	Disabled      bool          `json:"disabled,omitempty"` // Whether the market is disabled
}

type CrossMarkets struct {
	TradingPairs []V1CrossMarketsResult `json:"tradingPairs,omitempty"` // List of cross market trading pairs
}

type V1SpotMarketsResult struct {
	BaseIncrement  decimal.Decimal `json:"baseIncrement"`      // Min tick size increment of base currency
	Pair           *CurrencyPair   `json:"pair"`               // Currency pair information
	QuoteIncrement decimal.Decimal `json:"quoteIncrement"`     // Min tick price increment of quote currency or price
	Disabled       bool            `json:"disabled,omitempty"` // Whether the market is disabled
}

type SpotMarkets struct {
	TradingPairs []*V1SpotMarketsResult `json:"tradingPairs,omitempty"` // List of spot market trading pairs
}

type BlockchainNetwork struct {
	Coin                        string `json:"coin"`                        // Ticker symbol of native coin on chain
	MainnetBlockExplorerBaseUrl string `json:"mainnetBlockExplorerBaseUrl"` // Mainnet explorer URL
	MainnetName                 string `json:"mainnetName"`                 // Mainnet name
	TestnetBlockExplorerBaseUrl string `json:"testnetBlockExplorerBaseUrl"` // Testnet explorer URL
	TestnetName                 string `json:"testnetName"`                 // Testnet name
	Type                        string `json:"type"`                        // Chain ID
}

type TokenConfig struct {
	AssetType         string          `json:"assetType"`                   // Type of currency (e.g., "bridged")
	BridgeInfoUrl     string          `json:"bridgeInfoUrl,omitempty"`     // URL for more coin data
	CoinGeckoCurrency string          `json:"coinGeckoCurrency,omitempty"` // Ticker symbol on Coin Gecko
	CoinGeckoId       string          `json:"coinGeckoId,omitempty"`       // Coin ID on Coin Gecko
	Decimals          int             `json:"decimals"`                    // Max number of decimals for fractional values
	Description       string          `json:"description,omitempty"`       // Description of currency
	Id                string          `json:"id"`                          // Currency symbol
	MaxOrderSize      decimal.Decimal `json:"maxOrderSize,omitempty"`      // Maximum allowed order size
	MinOrderSize      decimal.Decimal `json:"minOrderSize,omitempty"`      // Minimum allowed order size
	Name              string          `json:"name"`                        // Currency name
	NativeAssetName   string          `json:"nativeAssetName,omitempty"`   // Native coin for wrapped coins
	Network           string          `json:"network"`                     // Blockchain network the coin is traded on
}

type Market struct {
	BlockchainNetwork []*BlockchainNetwork `json:"blockchainNetwork,omitempty"` // List of blockchain networks
	Cross             *CrossMarkets        `json:"cross,omitempty"`             // Cross market configuration
	Spot              *SpotMarkets         `json:"spot,omitempty"`              // Spot market configuration
	TokenConfig       []*TokenConfig       `json:"tokenConfig,omitempty"`       // List of allowed tokens by priority
}
