package cart

type CartItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}
