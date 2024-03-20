package main

import (
	"api/config"
	"api/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow requests from any origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false, // Allow cookies to be sent with requests
		MaxAge:           300,   // Maximum age for preflight requests
	}))

	// Routes definition
	router.Get("/", handlers.HealthCheckHandler)
	router.Get("/indexing", handlers.IndexHandler)
	router.Get("/inbox", handlers.InboxHandler)
	router.Get("/sent", handlers.SentHandler)

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
