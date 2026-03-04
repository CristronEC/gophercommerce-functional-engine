package catalog

import (
	"sync"

	"gophercommerce-functional-engine/internal/shared"
)

type service struct {
	products map[string]Product
	mu       sync.RWMutex
}

func NewService() Service {
	s := &service{
		products: make(map[string]Product),
	}

	// Seed inicial
	s.products["1"] = Product{"1", "Laptop", 1200}
	s.products["2"] = Product{"2", "Mouse", 25}
	s.products["3"] = Product{"3", "Keyboard", 50}

	return s
}

func (s *service) ListProducts() []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []Product
	for _, p := range s.products {
		list = append(list, p)
	}
	return list
}

func (s *service) GetProductByID(id string) (*Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[id]
	if !ok {
		return nil, shared.NewAppError("product not found", 404)
	}

	return &p, nil
}
