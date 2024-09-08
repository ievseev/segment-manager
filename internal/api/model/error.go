package model

import (
	"encoding/json"
	"net/http"
)

type ResultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResultError(code int, message string) *ResultError {
	return &ResultError{
		Code:    code,
		Message: message,
	}
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(NewResultError(statusCode, message))
}
