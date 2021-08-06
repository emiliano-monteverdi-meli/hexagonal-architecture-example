package cardsrv

import (
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type service struct {
	repository ports.CardRepository
}

func New(repository ports.CardRepository) ports.CardService {
	return &service{repository: repository}
}

func (srv *service) Create(card domain.Card) (domain.Card, error) {

	//
	// Validaciones
	//

	// ...

	if err := srv.repository.Save(card); err != nil {
		return domain.Card{}, err
	}

	return card, nil
}

func (srv *service) Get(id string) (domain.Card, error) {

	//
	// Validaciones
	//

	// ...

	card, err := srv.repository.Get(id)
	if err != nil {
		return domain.Card{}, err
	}

	return card, nil
}
