package transactionrepo

import (
	"errors"
	"fmt"

	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
)

func (repo *repository) Save(transaction domain.Transaction) error {

	//
	// Makes the KVS item
	//

	// ADAPTADOR // Golang (struct Transaction) ---> Item (KVS)
	transactionID := fmt.Sprintf("%v", transaction.ID)
	item := gokvsclient.MakeItem(transactionID, transaction)

	//
	// Saves the KVS item
	//

	if err := repo.kvsClient.Save(item); err != nil {
		return errors.New("saving the transaction into the kvs has failed")
	}

	return nil
}

func (repo *repository) Update(transaction domain.Transaction) error {

	//
	// Makes the KVS item
	//

	// ADAPTADOR // Golang (struct Transaction) ---> Item (KVS)
	transactionID := fmt.Sprintf("%v", transaction.ID)
	item := gokvsclient.MakeItem(transactionID, transaction)

	//
	// Updates the KVS item
	//

	if err := repo.kvsClient.Update(item); err != nil {
		return errors.New("updating the transaction into the kvs has failed")
	}

	return nil
}

func (repo *repository) Get(id string) (domain.Transaction, error) {

	//
	// Gets the KVS item
	//

	// ADAPTADOR // Golang (string id) ---> Item (KVS)
	item, err := repo.kvsClient.Get(id)
	if err != nil {
		return domain.Transaction{}, errors.New("getting the transaction from the kvs has failed")
	}

	if item == nil {
		return domain.Transaction{}, nil
	}

	transaction := domain.Transaction{}

	// ADAPTADOR // Item (KVS) ---> Golang (struct Transaction)
	if err := item.GetValue(&transaction); err != nil {
		return domain.Transaction{}, errors.New("item could not be marshaled into struct")
	}

	return transaction, nil
}
