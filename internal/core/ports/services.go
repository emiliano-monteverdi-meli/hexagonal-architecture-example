package ports

import "github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"

//go:generate mockgen -source=./services.go -package=mockups -destination=../../../mocks/mockups/services.go

type CardService interface {
	Create(card domain.Card) (domain.Card, error)
	Get(id string) (domain.Card, error)
}

type TransactionService interface {
	Create(transaction domain.Transaction) (domain.Transaction, error)
	Get(id string) (domain.Transaction, error)
}
