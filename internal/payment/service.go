package payment

type Service interface {
	Process(amount float64) PaymentResponse
}
