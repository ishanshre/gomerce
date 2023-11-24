package middleware

import "context"

type Middleware interface{}

type middleware struct {
	ctx context.Context
}

func NewMiddleware(ctx context.Context) Middleware {
	return &middleware{
		ctx: ctx,
	}
}
