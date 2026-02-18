package cart

// CartService define el contrato del módulo carrito.
// Se utiliza una interfaz para desacoplar implementación y lógica.
type CartService interface {
	AddItem(items []CartItem, item CartItem) []CartItem
	RemoveItem(items []CartItem, productID string) []CartItem
	CalculateTotal(items []CartItem) (float64, error)
}
