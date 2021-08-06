package dependencies

import (
	"net/http"
	"time"

	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest/retry"
)

type config struct {
	kvs map[string]kvs
	ds  map[string]ds
}

type kvs struct {
	name                    string
	readMaxIdleConnections  int
	writeMaxIdleConnections int
	readTimeout             time.Duration
	writeTimeout            time.Duration //nolint
	readRetryStrategy       map[string]int
	writeRetryStrategy      map[string]int
}

type ds struct {
	readMaxIdleConnections  int
	writeMaxIdleConnections int
	readTimeout             time.Duration
	writeTimeout            time.Duration
	retryStrategy           map[string]int
	readEndpoint            string
	writeEndpoint           string
	namespace               string
	entity                  string
}

var environmentConfigs = map[string]config{
	"sandbox": {
		kvs: map[string]kvs{
			"transaction-repository": {
				name:                    "TRANSACTION_STG",
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             50,
				writeTimeout:            50,
				readRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
				writeRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
			},
			"card-repository": {
				name:                    "CARD_STG",
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             50,
				writeTimeout:            50,
				readRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
				writeRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
			},
		},
		ds: map[string]ds{
			"transaction-repository": {
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             3000,
				writeTimeout:            3000,
				retryStrategy: map[string]int{
					"max_retries": 3,
					"time":        5,
				},
				namespace: "sandbox",
				entity:    "transaction-ds",
			},
		},
	},
	"production": {
		kvs: map[string]kvs{
			"transaction-repository": {
				name:                    "TRANSACTION_PROD",
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             50,
				writeTimeout:            50,
				readRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
				writeRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
			},
			"card-repository": {
				name:                    "CARD_PROD",
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             50,
				writeTimeout:            50,
				readRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
				writeRetryStrategy: map[string]int{
					"max_retries": 1,
					"time":        5,
				},
			},
		},
		ds: map[string]ds{
			"transaction-repository": {
				readMaxIdleConnections:  50,
				writeMaxIdleConnections: 50,
				readTimeout:             3000,
				writeTimeout:            3000,
				retryStrategy: map[string]int{
					"max_retries": 3,
					"time":        5,
				},
				namespace: "production",
				entity:    "transaction-ds",
			},
		},
	},
}

//
// KVS
//

func initKvs(config kvs) gokvsclient.Client {
	kvsConfig := gokvsclient.MakeKvsConfig()

	kvsConfig.SetReadMaxIdleConnections(config.readMaxIdleConnections)
	kvsConfig.SetWriteMaxIdleConnections(config.writeMaxIdleConnections)
	kvsConfig.SetReadTimeout(config.readTimeout * time.Millisecond)
	kvsConfig.SetWriteTimeout(config.readTimeout * time.Millisecond)
	kvsConfig.SetReadRetryStrategy(retry.NewSimpleRetryStrategy(config.readRetryStrategy["max_retries"], time.Duration(config.readRetryStrategy["time"])*time.Millisecond))
	kvsConfig.SetWriteRetryStrategy(retry.NewSimpleRetryStrategy(config.writeRetryStrategy["max_retries"], time.Duration(config.writeRetryStrategy["time"])*time.Millisecond, http.MethodPost, http.MethodPut, http.MethodGet))

	return gokvsclient.MakeKvsClient(config.name, kvsConfig)
}

//
// DS
//

func initDs(config ds) godsclient.Client {
	dsConfig := godsclient.NewDsClientConfig()

	dsConfig.WithReadMaxIdleConnections(config.readMaxIdleConnections)
	dsConfig.WithWriteMaxIdleConnections(config.writeMaxIdleConnections)
	dsConfig.WithReadTimeout(config.readTimeout * time.Millisecond)
	dsConfig.WithWriteTimeout(config.writeTimeout * time.Millisecond)
	dsConfig.WithRetryStrategy(retry.NewSimpleRetryStrategy(config.retryStrategy["max_retries"], time.Duration(config.retryStrategy["time"])*time.Millisecond))
	dsConfig.WithReadEndpoint(config.readEndpoint)
	dsConfig.WithWriteEndpoint(config.writeEndpoint)
	dsConfig.WithNamespace(config.namespace)
	dsConfig.WithEntity(config.entity)

	return godsclient.NewEntityClient(dsConfig)
}
