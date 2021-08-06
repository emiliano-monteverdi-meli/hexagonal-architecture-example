package cardrepo

import (
	"github.com/mercadolibre/hexagonal-architecture-example/internal/core/ports"
)

type memoryRepository struct {
	memory map[string]interface{}
}

func NewMemory() ports.CardRepository {
	return &memoryRepository{memory: make(map[string]interface{})}
}
