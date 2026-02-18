package catalog

// CatalogService define el contrato del módulo catálogo.
type CatalogService interface {
	GetAll() []Product
	FindByID(id string) (Product, error)
	AddProduct(products []Product, product Product) []Product
}
