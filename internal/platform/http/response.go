package http

import (
	"encoding/json"
	"estatehub-api/internal/platform/shared"
	"log"
	nethttp "net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteJson(w nethttp.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("error encoding json:", err)
	}
}

func ErrorJson(w nethttp.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Message: message})
	if err != nil {
		log.Println("error encoding json:", err)
	}

}

func ErrorJsonV2(w nethttp.ResponseWriter, statusCode int, err *shared.APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errEncode := json.NewEncoder(w).Encode(err)
	if errEncode != nil {
		log.Println("error encoding json:", errEncode)
	}

}
