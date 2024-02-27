package main

import (
	"api/config"
	"api/src"
	"log"
	"net/http"
)

func main() {
	port := config.GetEnv("API_PORT", "8080")

	// Initialize the router and set up routes
	router := src.Router()

	// Start the server
	log.Printf("Server listening on :%s...\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
