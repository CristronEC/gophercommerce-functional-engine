package shared

// JSONResponse es una estructura estándar para respuestas HTTP.
type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
