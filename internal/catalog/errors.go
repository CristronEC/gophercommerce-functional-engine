package catalog

import "errors"

var (
	ErrInvalidProductID = errors.New("invalid product ID")
	ErrInvalidPrice     = errors.New("invalid price")
	ErrInvalidStock     = errors.New("invalid stock value")
	ErrProductNotFound  = errors.New("product not found")
)
