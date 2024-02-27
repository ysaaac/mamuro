package src

import (
	"api/src/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Router creates and returns a new chi Router with defined routes.
func Router() http.Handler {
	router := chi.NewRouter()

	// Routes definition
	router.Get("/", handlers.HealthCheckHandler)

	return router
}
