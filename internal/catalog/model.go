package catalog

// Product representa un producto del sistema.
// Se aplica encapsulación manteniendo campos privados.
type Product struct {
	id       string
	name     string
	price    float64
	stock    int
	category string
}

// Constructor con validaciones
func NewProduct(id, name string, price float64, stock int, category string) (Product, error) {

	if id == "" {
		return Product{}, ErrInvalidProductID
	}

	if price <= 0 {
		return Product{}, ErrInvalidPrice
	}

	if stock < 0 {
		return Product{}, ErrInvalidStock
	}

	return Product{
		id:       id,
		name:     name,
		price:    price,
		stock:    stock,
		category: category,
	}, nil
}

// Getters
func (p Product) ID() string       { return p.id }
func (p Product) Name() string     { return p.name }
func (p Product) Price() float64   { return p.price }
func (p Product) Stock() int       { return p.stock }
func (p Product) Category() string { return p.category }
