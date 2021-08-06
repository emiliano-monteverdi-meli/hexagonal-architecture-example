package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest/retry"
	"github.com/mercadolibre/hexagonal-architecture-example/cmd/dependencies"
	"github.com/mercadolibre/hexagonal-architecture-example/cmd/httpserver"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/services/cardsrv"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/handlers/cardhdl"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/repositories/cardrepo"
	"github.com/mercadolibre/hexagonal-architecture-example/mocks"
)

func main() {

	//
	// Dependencies Injection
	//

	// OPT 1
	d1 := dependencies.NewByEnvironment()

	// OPT 2
	cardRepoKvs, transactionRepoKvs, transactionRepoDs := dependenciesOpt2()
	d2 := dependencies.New(cardRepoKvs, transactionRepoKvs, transactionRepoDs)

	// OPT 3
	cardRepositoryKvs, cardService, cardHandler := dependenciesOp3()
	d3 := dependencies.Definition{
		TransactionRepository: nil,
		CardRepository:        cardRepositoryKvs,
		TransactionService:    nil,
		CardService:           cardService,
		TransactionHandler:    nil,
		CardHandler:           cardHandler,
	}

	fmt.Println(d1, d2, d3)

	httpserver.Start(d1)
}

//
// OPT 2
//

func dependenciesOpt2() (gokvsclient.Client, gokvsclient.Client, godsclient.Client) {

	cardRepositoryKvs := initKvs("card_repository")
	transactionRepositoryKvs := initKvs("transaction_repository")
	transactionRepositoryDs := initDs("transaction_repository")

	return cardRepositoryKvs, transactionRepositoryKvs, transactionRepositoryDs
}

//
// KVS
//

func initKvs(name string) gokvsclient.Client {
	kvsConfig := gokvsclient.MakeKvsConfig()

	kvsConfig.SetReadMaxIdleConnections(10)
	kvsConfig.SetWriteMaxIdleConnections(10)
	kvsConfig.SetReadTimeout(10 * time.Millisecond)
	kvsConfig.SetWriteTimeout(10 * time.Millisecond)
	kvsConfig.SetReadRetryStrategy(retry.NewSimpleRetryStrategy(3, time.Duration(300)*time.Millisecond))
	kvsConfig.SetWriteRetryStrategy(retry.NewSimpleRetryStrategy(3, time.Duration(300)*time.Millisecond, http.MethodPost, http.MethodPut, http.MethodGet))

	return gokvsclient.MakeKvsClient(name, kvsConfig)
}

//
// DS
//

func initDs(name string) godsclient.Client {
	dsConfig := godsclient.NewDsClientConfig()

	dsConfig.WithReadMaxIdleConnections(10)
	dsConfig.WithWriteMaxIdleConnections(10)
	dsConfig.WithReadTimeout(10 * time.Millisecond)
	dsConfig.WithWriteTimeout(10 * time.Millisecond)
	dsConfig.WithRetryStrategy(retry.NewSimpleRetryStrategy(3, time.Duration(300)*time.Millisecond))
	dsConfig.WithEntity(name)

	return godsclient.NewEntityClient(dsConfig)
}

//
// OPT 3
//

func dependenciesOp3() (ports.CardRepository, ports.CardService, *cardhdl.Handler) {

	//
	// Repositories
	//

	cardRepository := cardrepo.NewKVS(mocks.NewKvsClient())

	//
	// Core
	//

	cardService := cardsrv.New(cardRepository)

	//
	// Handlers
	//

	cardHandler := cardhdl.New(cardService)

	return cardRepository, cardService, cardHandler

}
