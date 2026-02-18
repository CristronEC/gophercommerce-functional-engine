package cart

import "github.com/samber/lo"

// cartServiceImpl es la implementación concreta del contrato CartService.
// No mantiene estado interno, lo que lo hace funcional y testeable.
type cartServiceImpl struct{}

// NewCartService crea una nueva instancia del servicio.
func NewCartService() CartService {
	return &cartServiceImpl{}
}

// AddItem retorna un nuevo slice sin mutar el original.
func (s *cartServiceImpl) AddItem(items []CartItem, item CartItem) []CartItem {
	return append(items, item)
}

// RemoveItem filtra el producto sin modificar el slice original.
func (s *cartServiceImpl) RemoveItem(items []CartItem, productID string) []CartItem {
	return lo.Filter(items, func(item CartItem, _ int) bool {
		return item.ProductID() != productID
	})
}

// CalculateTotal aplica reducción funcional sobre el carrito.
func (s *cartServiceImpl) CalculateTotal(items []CartItem) (float64, error) {

	if len(items) == 0 {
		return 0, ErrEmptyCart
	}

	total := lo.Reduce(items, func(acc float64, item CartItem, _ int) float64 {
		return acc + (item.Price() * float64(item.Quantity()))
	}, 0.0)

	return total, nil
}
