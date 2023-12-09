package middleware

import (
	"context"
	"net/http"

	"github.com/ishanshre/gomerce/internals/config"
	"github.com/justinas/nosurf"
)

type Middleware interface {
	SessionLoad(next http.Handler) http.Handler
	NoSurf(next http.Handler) http.Handler
}

type middleware struct {
	app *config.AppConfig
	ctx context.Context
}

func NewMiddleware(app *config.AppConfig, ctx context.Context) Middleware {
	return &middleware{
		app: app,
		ctx: ctx,
	}
}

// SessionLoad loads and saves the session on every request
func (m *middleware) SessionLoad(next http.Handler) http.Handler {
	return m.app.Session.LoadAndSave(next)
}

// NoSurf implement csrf token middleware
func (m *middleware) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next) // creates a new handler
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   m.app.InProduction,
		SameSite: http.SameSiteLaxMode, // allows cookies to sent in cross site
	})
	return csrfHandler
}
