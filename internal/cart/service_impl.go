package cart

import "sync"

type service struct {
	items []CartItem
	mu    sync.RWMutex
}

func NewService() Service {
	return &service{
		items: []CartItem{},
	}
}

func (s *service) AddItem(productID string, name string, price float64, quantity int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item := CartItem{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
	}

	s.items = append(s.items, item)
	return nil
}

func (s *service) GetCart() Cart {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var total float64
	for _, i := range s.items {
		total += i.Price * float64(i.Quantity)
	}

	return Cart{
		Items: s.items,
		Total: total,
	}
}

func (s *service) ClearCart() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = []CartItem{}
}
