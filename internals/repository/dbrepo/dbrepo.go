package dbrepo

import (
	"context"

	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/connection"
	"github.com/ishanshre/gomerce/internals/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  connection.Connection
	Ctx context.Context
}

func NewPostgresRepo(conn connection.Connection, a *config.AppConfig, ctx context.Context) repository.Repository {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
		Ctx: ctx,
	}
}
