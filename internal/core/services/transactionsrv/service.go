package transactionsrv

import (
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type service struct {
	repository ports.TransactionRepository
}

func New(repository ports.TransactionRepository) ports.TransactionService {
	return &service{repository: repository}
}

func (srv *service) Create(transaction domain.Transaction) (domain.Transaction, error) {

	//
	// Validaciones
	//

	// ...

	if err := srv.repository.Save(transaction); err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (srv *service) Get(id string) (domain.Transaction, error) {

	//
	// Validaciones
	//

	// ...

	transaction, err := srv.repository.Get(id)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}
