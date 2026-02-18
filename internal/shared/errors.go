package shared

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrInvalidJSON    = errors.New("invalid JSON payload")
)
