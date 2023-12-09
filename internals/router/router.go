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
	router.Use(m.SessionLoad) // load the session middleware
	router.Use(m.NoSurf)      // csrf middleware
	router.Use(chi_middlewares.Logger)

	static_file_server := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", static_file_server))

	router.Route("/api/v1", func(router chi.Router) {
		router.Post("/category", h.PostCategoryHandler)
		router.Get("/category", h.GetCategoriesHandler)
		router.Get("/category/{id}", h.GetCategoryHandler)
		router.Delete("/category/{id}", h.DeleteCategoryHandler)
		router.Put("/category/{id}", h.UpdateCategoryHandler)
	})

	router.Get("/", h.HomePageHandler)
	return router
}
