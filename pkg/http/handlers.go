package http

import (
	"encoding/json"
	"net/http"

	"gophercommerce-functional-engine/internal/catalog"
	"gophercommerce-functional-engine/internal/shared"
)

type Handler struct {
	catalogService catalog.CatalogService
}

func NewHandler(catalogService catalog.CatalogService) *Handler {
	return &Handler{catalogService: catalogService}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := shared.JSONResponse{
		Status:  "success",
		Message: "API is running",
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	products := h.catalogService.GetAll()

	response := shared.JSONResponse{
		Status:  "success",
		Message: "Products retrieved",
		Data:    products,
	}

	json.NewEncoder(w).Encode(response)
}
