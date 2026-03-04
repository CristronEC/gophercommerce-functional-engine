package cart

type Service interface {
	AddItem(productID string, name string, price float64, quantity int) error
	GetCart() Cart
	ClearCart()
}
