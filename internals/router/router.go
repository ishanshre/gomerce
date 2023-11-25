package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chi_middlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/handler"
	"github.com/ishanshre/gomerce/internals/middleware"
)

func Router(app *config.AppConfig, h handler.Handler, m middleware.Middleware) http.Handler {
	router := chi.NewRouter()
	router.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})))
	router.Use(chi_middlewares.Logger)

	router.Route("/api/v1", func(router chi.Router) {
		router.Post("/category", h.PostCategoryHandler)
	})
	return router
}
