package shared

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func HandleError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*AppError); ok {
		WriteJSON(w, appErr.Code, JSONResponse{
			Success: false,
			Error:   appErr.Message,
		})
		return
	}

	WriteJSON(w, http.StatusInternalServerError, JSONResponse{
		Success: false,
		Error:   "internal server error",
	})
}
