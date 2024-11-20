package api

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewError(w http.ResponseWriter, statusCode int, err error) {
	http.Error(w, err.Error(), statusCode)
}

func Result(w http.ResponseWriter, statusCode int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Message: msg,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
