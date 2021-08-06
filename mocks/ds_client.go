package mocks

import (
	"github.com/mercadolibre/go-meli-toolkit/godsclient"
)

func NewDsClient() *godsclient.MockEntityClient {
	return new(godsclient.MockEntityClient)
}
