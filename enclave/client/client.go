package client

type Client interface {
	SpotClient
	PerpsClient
}

type client struct {
	SpotClient
	PerpsClient
}

func NewClient(apiKey, apiSecret, baseURL string) Client {
	return NewClientWithBase(NewBaseClient(apiKey, apiSecret, baseURL))
}

func NewClientWithBase(base BaseClient) Client {
	return &client{
		SpotClient:  NewSpotClientWithBase(base),
		PerpsClient: NewPerpsClientWithBase(base),
	}
}
