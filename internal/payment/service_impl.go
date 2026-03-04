package payment

import "time"

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Process(amount float64) PaymentResponse {
	time.Sleep(1 * time.Second)

	return PaymentResponse{
		Status:  "approved",
		Amount:  amount,
		Message: "payment processed successfully",
	}
}
