package httpserver

import (
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/mercadolibre/hexagonal-architecture-example/cmd/dependencies"
)

func Routes(router *fury.Application, d dependencies.Definition) {

	//
	// Transactions
	//

	router.Post("/gateway/transactions", d.TransactionHandler.Create)
	router.Post("/gateway/transactions/:id", d.TransactionHandler.Get)

	//
	// Card
	//

	router.Post("/gateway/card", d.CardHandler.Create)
	router.Post("/gateway/card/:id", d.CardHandler.Get)
}
