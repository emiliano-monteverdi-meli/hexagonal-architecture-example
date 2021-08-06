package cardrepo

import (
	"errors"

	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
)

func (repo *memoryRepository) Save(card domain.Card) error {

	repo.memory[card.ID] = card

	return nil
}

func (repo *memoryRepository) Get(id string) (domain.Card, error) {

	value := repo.memory[id]

	if value == nil {
		return domain.Card{}, errors.New("refund id has nil value")
	}

	refund := value.(domain.Card)

	return refund, nil
}
