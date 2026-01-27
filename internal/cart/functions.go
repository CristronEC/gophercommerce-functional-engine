type CartItem struct {
	ProductID string
	Price     float64
	Quantity  int
}

func CalculateTotal[T any](items []T, mapper func(T) float64) float64 {
	return lo.Reduce(items, func(acc float64, item T, _ int) float64 {
		return acc + mapper(item)
	}, 0)
}

func CartTotal(items []CartItem) float64 {
	return CalculateTotal(items, func(item CartItem) float64 {
		return item.Price * float64(item.Quantity)
	})
}