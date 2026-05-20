package http

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse defines the expected JSON structure for API errors.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WriteJSON writes a standard JSON response with the given status code.
func WriteJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if payload != nil {
		_ = json.NewEncoder(w).Encode(payload)
	}
}

// WriteError writes a standardized JSON error response.
func WriteError(w http.ResponseWriter, statusCode int, code string, message string) {
	payload := ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
	WriteJSON(w, statusCode, payload)
}
