package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// RouterConfig contains the BusinessLogic for the router.
type RouterConfig struct {
	BusinessLogic Services
}

// Handlers contains all the interfaces for rest handlers
type Handlers struct {
	Services Services
}

// NewRouter returns a new Chi router with all the config, handlers,
// routes and options already set.
func NewRouter(cfg RouterConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// not production worthy - but this is just a local app for now
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	handlers := Handlers{
		Services: cfg.BusinessLogic,
	}

	r.Post("/url", handlers.CreateShortUrl)
	r.Get("/{url}", handlers.GetOriginalUrl)

	return r
}
