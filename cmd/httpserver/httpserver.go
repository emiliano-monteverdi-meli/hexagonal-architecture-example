package httpserver

import (
	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/mercadolibre/hexagonal-architecture-example/cmd/dependencies"
)

func Start(d dependencies.Definition) {

	if err := SetupRouter(d).Run(); err != nil {
		log.Errors("error running server", []error{err})
	}
}

func SetupRouter(d dependencies.Definition) *fury.Application {
	router, err := fury.NewWebApplication()
	if err != nil {
		log.Errors("can't initialize app", []error{err})
	}

	Routes(router, d)

	return router
}
