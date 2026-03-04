package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"gophercommerce-functional-engine/internal/cart"
	"gophercommerce-functional-engine/internal/catalog"
	"gophercommerce-functional-engine/internal/payment"
	"gophercommerce-functional-engine/internal/shared"
)

type Handler struct {
	catalog catalog.Service
	cart    cart.Service
	payment payment.Service
}

func NewHandler(c catalog.Service, ca cart.Service, p payment.Service) *Handler {
	return &Handler{
		catalog: c,
		cart:    ca,
		payment: p,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := h.catalog.ListProducts()
	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    products,
	})
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/catalog/")

	product, err := h.catalog.GetProductByID(id)
	if err != nil {
		shared.HandleError(w, err)
		return
	}

	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    product,
	})
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.HandleError(w, shared.NewAppError("invalid request", 400))
		return
	}

	product, err := h.catalog.GetProductByID(req.ProductID)
	if err != nil {
		shared.HandleError(w, err)
		return
	}

	h.cart.AddItem(product.ID, product.Name, product.Price, req.Quantity)

	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    "item added",
	})
}

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	cart := h.cart.GetCart()

	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    cart,
	})
}

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	h.cart.ClearCart()

	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    "cart cleared",
	})
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	cart := h.cart.GetCart()

	if cart.Total == 0 {
		shared.HandleError(w, shared.NewAppError("cart is empty", 400))
		return
	}

	resp := h.payment.Process(cart.Total)
	h.cart.ClearCart()

	shared.WriteJSON(w, 200, shared.JSONResponse{
		Success: true,
		Data:    resp,
	})
}
