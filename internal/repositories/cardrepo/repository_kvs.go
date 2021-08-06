package cardrepo

import (
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type kvsRepository struct {
	kvsClient gokvsclient.Client
}

func NewKVS(kvsClient gokvsclient.Client) ports.CardRepository {
	return &kvsRepository{kvsClient: kvsClient}
}
