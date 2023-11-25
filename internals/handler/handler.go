package handler

import (
	"context"

	validators "github.com/go-playground/validator"
	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/connection"
	"github.com/ishanshre/gomerce/internals/repository"
	"github.com/ishanshre/gomerce/internals/validator"
)

// interface for handler
type Handler interface{}

// handler struct
type handler struct {
	app  *config.AppConfig
	repo repository.Repository
	conn connection.Connection
	ctx  context.Context
}

// declare validator
var validate *validators.Validate

// intialize a handler
func NewHandler(app *config.AppConfig, repo repository.Repository, conn connection.Connection, ctx context.Context) Handler {
	validate = validators.New()
	validate.RegisterValidation("upper", validator.UpperCase)
	validate.RegisterValidation("lower", validator.LowerCase)
	validate.RegisterValidation("number", validator.Number)
	return &handler{
		app:  app,
		repo: repo,
		conn: conn,
		ctx:  ctx,
	}
}
