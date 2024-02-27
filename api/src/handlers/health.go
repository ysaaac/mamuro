package handlers

import (
	"net/http"
)

func HealthCheck() ResponseStructure {
	return ResponseStructure{
		Status:  "ok",
		Message: "Server is healthy",
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheck()
	WriteJsonResponse(w, http.StatusOK, response)
}
