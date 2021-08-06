package transactionrepo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/domain"
	"github.com/mercadolibre/hexagonal-architecture-example/internal/repositories/transactionrepo"
	"github.com/mercadolibre/hexagonal-architecture-example/mocks"
	"github.com/mercadolibre/hexagonal-architecture-example/mocks/mockups"
	"github.com/stretchr/testify/assert"
)

type dependencies struct {
	kvsClient *mockups.MockKvsClient
	dsClient  *godsclient.MockEntityClient
}

func makeDependencies(t *testing.T) dependencies {
	return dependencies{
		kvsClient: mockups.NewMockKvsClient(gomock.NewController(t)),
		dsClient:  mocks.NewDsClient(),
	}
}

func TestSave(t *testing.T) {

	//
	// Setup
	//

	transaction := mocks.Transaction()

	//
	// Errors
	//

	saveKvsError := errors.New("saving the transaction into the kvs has failed")

	//
	// Tests Cases
	//

	type args struct {
		transaction domain.Transaction
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m dependencies, args args, want want)
	}{
		{
			name: "Should create a transaction successfully",
			mock: func(m dependencies, args args, want want) {
				cardID := fmt.Sprintf("%v", args.transaction.ID)
				item := gokvsclient.MakeItem(cardID, args.transaction)

				m.kvsClient.EXPECT().Save(item).Return(nil)
			},
			args: args{
				transaction: transaction,
			},
			want: want{},
		},
		{
			name: "Should fail when tries to create a transaction",
			mock: func(d dependencies, args args, want want) {
				cardID := fmt.Sprintf("%v", args.transaction.ID)
				item := gokvsclient.MakeItem(cardID, args.transaction)

				d.kvsClient.EXPECT().Save(item).Return(errors.New("fail"))
			},
			args: args{
				transaction: transaction,
			},
			want: want{
				err: saveKvsError,
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
			d := makeDependencies(t)
			tt.mock(d, args, want)

			repository := transactionrepo.New(d.kvsClient, d.dsClient)

			//
			// Execute
			//

			err := repository.Save(args.transaction)

			//
			// Verify
			//

			if err != nil && want.err != nil {
				assert.Equal(t, want.err.Error(), err.Error())
			} else {
				assert.Equal(t, want.err, err)
			}
		})
	}
}
