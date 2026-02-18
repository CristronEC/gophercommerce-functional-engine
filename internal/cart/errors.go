package cart

import "errors"

var (
	ErrInvalidProductID = errors.New("invalid product ID")
	ErrInvalidPrice     = errors.New("invalid price")
	ErrInvalidQuantity  = errors.New("invalid quantity")
	ErrEmptyCart        = errors.New("cart is empty")
)
