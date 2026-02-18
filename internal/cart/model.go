package cart

// CartItem representa un producto dentro del carrito.
// Sus campos son privados para garantizar encapsulación.
type CartItem struct {
	productID string
	price     float64
	quantity  int
}

// NewCartItem es un constructor que valida datos antes de crear el objeto.
func NewCartItem(productID string, price float64, quantity int) (CartItem, error) {
	if productID == "" {
		return CartItem{}, ErrInvalidProductID
	}

	if price <= 0 {
		return CartItem{}, ErrInvalidPrice
	}

	if quantity <= 0 {
		return CartItem{}, ErrInvalidQuantity
	}

	return CartItem{
		productID: productID,
		price:     price,
		quantity:  quantity,
	}, nil
}

// Métodos getters (encapsulación)
func (c CartItem) ProductID() string {
	return c.productID
}

func (c CartItem) Price() float64 {
	return c.price
}

func (c CartItem) Quantity() int {
	return c.quantity
}
