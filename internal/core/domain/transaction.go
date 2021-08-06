package domain

import "github.com/mercadolibre/go-meli-toolkit/godsclient"

type Transaction struct {
	ID        string `json:"id"`
	Amount    string `json:"amount"`
	CreatedAt string `json:"created_at"`
}

// ejemplo extremadamente acotado, ver la implementacion completa en gateway-apicard-go
type TransactionResponse struct {
	ScrollResponse *godsclient.ScrollResponse
}
