package middleware

import (
	"context"

	"github.com/ishanshre/gomerce/internals/config"
)

type Middleware interface{}

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
