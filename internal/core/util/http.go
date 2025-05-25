package util

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/go-playground/validator/v10"

)

type ErrorResponse struct {
	Message string
}

var Validate = validator.New()

func SendSuccessResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {

	log.Println("error:", message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	json.NewEncoder(w).Encode(&ErrorResponse{Message: message})
}

