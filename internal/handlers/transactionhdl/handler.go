package transactionhdl

import (
	"encoding/json"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type Handler struct {
	service ports.TransactionService
}

func New(service ports.TransactionService) *Handler {
	return &Handler{service: service}
}

func (hdl *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	body := domain.Transaction{}

	// ADAPTADOR // HTTP ---> Golang (struct Transaction)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return web.NewError(400, "unmarshalling error")
	}

	card, err := hdl.service.Create(body)
	if err != nil {
		return web.NewError(500, "create service error")
	}

	// ADAPTADOR // Golang (struct Transaction) ---> HTTP
	return web.RespondJSON(w, card, http.StatusOK)
}

func (hdl *Handler) Get(w http.ResponseWriter, r *http.Request) error {

	// ADAPTADOR // HTTP ---> Golang (string)
	transactionID := web.Params(r)["id"]

	card, err := hdl.service.Get(transactionID)
	if err != nil {
		return web.NewError(500, "get service error")
	}

	// ADAPTADOR // Golang (struct Transaction) ---> HTTP
	return web.RespondJSON(w, card, http.StatusOK)
}
