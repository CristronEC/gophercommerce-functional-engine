package payment

type PaymentRequest struct {
	Amount float64 `json:"amount"`
}

type PaymentResponse struct {
	Status  string  `json:"status"`
	Amount  float64 `json:"amount"`
	Message string  `json:"message"`
}
