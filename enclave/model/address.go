package model

type Address struct {
	Address string `json:"address"` // Address provisioned for a specific coin
	Coin    string `json:"coin"`    // Coin for which the customer provisioned an address
}

type ProvisionedAddress struct {
	AccountId string `json:"accountId"` // Internal ID associated with the account that made the request
	Address   string `json:"address"`   // The provisioned deposit address
	Symbol    string `json:"symbol"`    // Symbol associated with the deposit (encapsulated subtype)
}
