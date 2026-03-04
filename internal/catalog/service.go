package catalog

type Service interface {
	ListProducts() []Product
	GetProductByID(id string) (*Product, error)
}
