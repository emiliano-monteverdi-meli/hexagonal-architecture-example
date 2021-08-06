package mocks

import "github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"

func Card() domain.Card {
	return domain.Card{
		ID:        "123",
		UserID:    "9000000",
		CreatedAt: "10/09/2077",
	}
}
