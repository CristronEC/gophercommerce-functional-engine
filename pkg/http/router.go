package http

import (
	"net/http"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/catalog", handler.ListProducts)
	mux.HandleFunc("/catalog/", handler.GetProduct)

	mux.HandleFunc("/cart", handler.GetCart)
	mux.HandleFunc("/cart/add", handler.AddToCart)
	mux.HandleFunc("/cart/clear", handler.ClearCart)

	mux.HandleFunc("/checkout", handler.Checkout)

	return mux
}
