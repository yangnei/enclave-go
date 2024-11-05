package client

type SpotClient interface {
	OrderFillClient
}

// NewSpotClientWithBase initializes a new orderFillClient client with the provided baseClient.
func NewSpotClientWithBase(baseClient BaseClient) SpotClient {
	return &orderFillClient{
		BaseClient: baseClient,
		prefix:     "/v1",
	}
}

// NewSpotClient initializes a new orderFillClient client with the provided API key and secret.
func NewSpotClient(apiKey, apiSecret, baseURL string) SpotClient {
	return NewSpotClientWithBase(NewBaseClient(apiKey, apiSecret, baseURL))
}
