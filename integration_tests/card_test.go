package integration_tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mercadolibre/hexagonal-architecture-example/cmd/dependencies"

	"github.com/mercadolibre/hexagonal-architecture-example/cmd/httpserver"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/services/cardsrv"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/handlers/cardhdl"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/repositories/cardrepo"
	"github.com/mercadolibre/hexagonal-architecture-example/mocks"
	"github.com/stretchr/testify/assert"
)

type cardDependencies struct {
	cardRepository ports.CardRepository
	cardService    ports.CardService
	cardHandler    *cardhdl.Handler
}

func makeCardDependencies() cardDependencies {

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

	return cardDependencies{cardRepository: cardRepository, cardService: cardService, cardHandler: cardHandler}
}

func TestSave(t *testing.T) {

	//
	// Setup
	//

	card := mocks.Card()
	cardBody, _ := json.Marshal(card)

	//
	// Tests Cases
	//

	type args struct {
		url     string
		headers map[string]string
		body    []byte
	}

	type want struct {
		status int
		result interface{}
		error  error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should create the card v4 successfully",
			args: args{
				url:     "/gateway/card",
				headers: map[string]string{"X-Request-ID": "123-456-789"},
				body:    cardBody,
			},
			want: want{
				status: 200,
				result: card,
			},
		},
	}

	//
	// Tests Runner
	//

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			//
			// Setup
			//

			args, want := tt.args, tt.want

			m := makeCardDependencies()

			//
			// Execute
			//

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", args.url, strings.NewReader(string(args.body)))
			for k, v := range args.headers {
				req.Header.Set(k, v)
			}

			//
			// Dependency Injection
			//

			d := dependencies.Definition{
				TransactionRepository: nil,
				CardRepository:        m.cardRepository,
				TransactionService:    nil,
				CardService:           m.cardService,
				TransactionHandler:    nil,
				CardHandler:           m.cardHandler,
			}

			httpserver.SetupRouter(d).ServeHTTP(w, req)

			//
			// Verify
			//

			assert.Equal(t, want.status, w.Code)
			fmt.Println(w.Body.String())
			if w.Code >= 200 && w.Code <= 299 {
				want, _ := json.Marshal(want.result)
				assert.Equal(t, string(want), w.Body.String())

			} else {
				want, _ := json.Marshal(want.error)
				assert.Equal(t, string(want), w.Body.String())
			}
		})
	}
}
