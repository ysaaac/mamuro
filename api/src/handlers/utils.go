package handlers

import (
	"encoding/json"
	"net/http"
)

// ResponseStructure represents the structure of the response JSON.
type ResponseStructure struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	// Set content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode response as JSON and write to the response writer
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
}
