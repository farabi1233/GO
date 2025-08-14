package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode) // Set response status to Created
	encoder := json.NewEncoder(w)
	encoder.Encode(data) // Encode the new product as JSON and write to response
}

func DD(v interface{}) {
	fmt.Printf("%+v\n", v)
	os.Exit(1) // stops program like dd()
}
