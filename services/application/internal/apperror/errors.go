package apperror

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewBadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Message:    message,
	}
}

func NewNotFound(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Message:    message,
	}
}

func NewInternal(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_ERROR",
		Message:    message,
	}
}

func NewValidation(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Code:       "VALIDATION_ERROR",
		Message:    message,
	}
}

// Write sends the error response as JSON
func Write(w http.ResponseWriter, err *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}
