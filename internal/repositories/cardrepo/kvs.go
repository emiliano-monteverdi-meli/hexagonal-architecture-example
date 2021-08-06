package cardrepo

import (
	"errors"

	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
)

func (repo *kvsRepository) Save(card domain.Card) error {

	//
	// Makes the KVS item
	//

	// ADAPTADOR // Golang (struct Card) ---> Item (KVS)
	item := gokvsclient.MakeItem(card.ID, card)

	//
	// Saves the KVS item
	//

	if err := repo.kvsClient.Save(item); err != nil {
		return errors.New("save card into kvs has fail")
	}

	return nil
}

func (repo *kvsRepository) Get(id string) (domain.Card, error) {

	//
	// Gets the KVS item
	//

	// ADAPTADOR // Golang (string id) ---> Item (KVS)
	item, err := repo.kvsClient.Get(id)
	if err != nil {
		return domain.Card{}, errors.New("get card from kvs has fail")
	}

	if item == nil {
		return domain.Card{}, nil
	}

	card := domain.Card{}

	// ADAPTADOR // Item (KVS) ---> Golang (struct Card)
	if err := item.GetValue(&card); err != nil {
		return domain.Card{}, errors.New("item could not be marshaled into struct")
	}

	return card, nil
}
