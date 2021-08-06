package dependencies

import (
	"os"

	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/services/cardsrv"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/services/transactionsrv"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/handlers/cardhdl"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/handlers/transactionhdl"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/repositories/cardrepo"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/repositories/transactionrepo"
	"github.com/mercadolibre/hexagonal-architecture-example/mocks"
)

type Definition struct {

	//
	// Repositories (Storage)
	//

	TransactionRepository ports.TransactionRepository
	CardRepository        ports.CardRepository

	//
	// Core (Business Logic)
	//

	TransactionService ports.TransactionService
	CardService        ports.CardService

	//
	// Handlers (Controller)
	//

	TransactionHandler *transactionhdl.Handler
	CardHandler        *cardhdl.Handler
}

func New(cardRepoKvs gokvsclient.Client, transactionRepoKvs gokvsclient.Client, transactionRepoDs godsclient.Client) Definition {

	d := initDependencies(cardRepoKvs, transactionRepoKvs, transactionRepoDs)

	return d
}

func NewByEnvironment() Definition {
	var environment string

	if os.Getenv("GO_ENVIRONMENT") == "production" {
		isTestScope := os.Getenv("FURY_IS_TEST_SCOPE")
		if isTestScope == "false" {
			environment = "production"
		} else {
			environment = "sandbox"
		}
	} else {
		d := initDependencies(mocks.NewKvsClient(), mocks.NewKvsClient(), mocks.NewDsClient())
		return d
	}

	configs := environmentConfigs[environment]

	// KVS
	cardRepositoryKvs := initKvs(configs.kvs["card-repository"])
	transactionRepositoryKvs := initKvs(configs.kvs["transaction-repository"])

	// DS
	transactionRepositoryDs := initDs(configs.ds["transaction-repository"])

	d := initDependencies(cardRepositoryKvs, transactionRepositoryKvs, transactionRepositoryDs)

	return d
}

func initDependencies(cardRepositoryKvs gokvsclient.Client, transactionRepositoryKvs gokvsclient.Client, transactionRepositoryDs godsclient.Client) Definition {
	d := Definition{}

	//
	// Repositories (Storage)
	//

	d.TransactionRepository = transactionrepo.New(transactionRepositoryKvs, transactionRepositoryDs)
	d.CardRepository = cardrepo.NewKVS(cardRepositoryKvs)

	//
	// Core (Business Logic)
	//

	d.TransactionService = transactionsrv.New(d.TransactionRepository)
	d.CardService = cardsrv.New(d.CardRepository)

	//
	// Handler (Controller)
	//

	d.TransactionHandler = transactionhdl.New(d.TransactionService)
	d.CardHandler = cardhdl.New(d.CardService)

	return d
}
