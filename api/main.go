package main

import (
	"api/config"
	"api/src/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	// Routes definition
	router.Get("/", handlers.HealthCheckHandler)
	router.Get("/indexing", handlers.IndexHandler)

	return router
}

func main() {
	port := config.GetEnv("API_PORT", "8080")

	// Initialize the router and set up routes
	router := Router()

	// Start the server
	log.Printf("Server listening on :%s...\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
