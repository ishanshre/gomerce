package handler

import (
	"context"

	"github.com/ishanshre/gomerce/internals/connection"
	"github.com/ishanshre/gomerce/internals/repository"
)

// interface for handler
type Handler interface{}

// handler struct
type handler struct {
	repo repository.Repository
	conn connection.Connection
	ctx  context.Context
}

// intialize a handler
func NewHandler(repo repository.Repository, conn connection.Connection, ctx context.Context) Handler {
	return &handler{
		repo: repo,
		conn: conn,
		ctx:  ctx,
	}
}
