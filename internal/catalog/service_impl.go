package catalog

import "github.com/samber/lo"

type catalogServiceImpl struct {
	products []Product
}

// Constructor con productos iniciales en memoria
func NewCatalogService() CatalogService {
	return &catalogServiceImpl{
		products: []Product{},
	}
}

func (s *catalogServiceImpl) GetAll() []Product {
	return s.products
}

func (s *catalogServiceImpl) FindByID(id string) (Product, error) {

	product, found := lo.Find(s.products, func(p Product) bool {
		return p.ID() == id
	})

	if !found {
		return Product{}, ErrProductNotFound
	}

	return product, nil
}

func (s *catalogServiceImpl) AddProduct(products []Product, product Product) []Product {
	return append(products, product)
}
