package dbrepo

import (
	"context"
	"database/sql"

	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
	Ctx context.Context
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig, ctx context.Context) repository.Repository {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
		Ctx: ctx,
	}
}
