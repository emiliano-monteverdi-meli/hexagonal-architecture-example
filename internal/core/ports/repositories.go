package ports

import "github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"

//go:generate mockgen -source=./repositories.go -package=mockups -destination=../../../mocks/mockups/repositories.go

type CardRepository interface {
	Save(card domain.Card) error
	Get(id string) (domain.Card, error)
}

type TransactionRepository interface {
	Save(transaction domain.Transaction) error
	Update(transaction domain.Transaction) error
	Get(id string) (domain.Transaction, error)

	FindAllExpired() (domain.TransactionResponse, error)
}
