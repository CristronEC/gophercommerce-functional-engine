package main

import (
	"log"
	"net/http"

	"gophercommerce-functional-engine/internal/cart"
	"gophercommerce-functional-engine/internal/catalog"
	"gophercommerce-functional-engine/internal/payment"
	httpHandler "gophercommerce-functional-engine/pkg/http"
)

func main() {
	catalogService := catalog.NewService()
	cartService := cart.NewService()
	paymentService := payment.NewService()

	handler := httpHandler.NewHandler(catalogService, cartService, paymentService)
	router := httpHandler.NewRouter(handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
