package transactionrepo

import (
	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type repository struct {
	kvsClient gokvsclient.Client
	dsClient  godsclient.Client
}

func New(kvsClient gokvsclient.Client, dsClient godsclient.Client) ports.TransactionRepository {
	return &repository{kvsClient: kvsClient, dsClient: dsClient}
}
