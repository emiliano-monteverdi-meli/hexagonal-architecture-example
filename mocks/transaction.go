package mocks

import "github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"

func Transaction() domain.Transaction {
	return domain.Transaction{
		ID:        "123",
		Amount:    "1000000",
		CreatedAt: "10/09/2077",
	}
}
