package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode) // Set response status to Created
	encoder := json.NewEncoder(w)
	encoder.Encode(data) // Encode the new product as JSON and write to response
}

func SendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  statusCode,
		"message": message,
		"data":    nil,
	})
}
