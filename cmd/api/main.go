package main

import (
	"log"
	"net/http"

	"gophercommerce-functional-engine/internal/catalog"
	httpPkg "gophercommerce-functional-engine/pkg/http"
)

func main() {

	catalogService := catalog.NewCatalogService()

	handler := httpPkg.NewHandler(catalogService)

	router := httpPkg.NewRouter(handler)

	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
