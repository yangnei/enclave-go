package model

type Hello struct {
	Hello string `json:"hello"`
}

type AuthenticatedHello struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}
