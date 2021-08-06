package transactionrepo

import (
	"errors"

	"github.com/mercadolibre/go-meli-toolkit/godsclient/querybuilders"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
)

const (
	pageSize      = 500
	statusExpired = "expired"
)

func (repo *repository) FindAllExpired() (domain.TransactionResponse, error) {

	//
	// Builds the ds query: range values don't include the actual number, so the actual 'from value' is the <number plus one> and the 'until value' is the <number minus one>
	//

	query := querybuilders.And(
		querybuilders.Eq("expiration_year", 2019),
		querybuilders.Eq("expiration_month", 1),
		querybuilders.Not(
			querybuilders.Eq("status", statusExpired),
		),
	)

	//
	// Executes the query and gets the objects
	//

	response, err := repo.dsClient.ScrollBuilder().WithSize(int32(pageSize)).WithQuery(query).IsSecondarySearch(true).AddProjection("id").Execute()
	if err != nil {
		return domain.TransactionResponse{}, errors.New("getting the expired cards from the ds has failed")
	}

	return domain.TransactionResponse{ScrollResponse: response}, nil
}
