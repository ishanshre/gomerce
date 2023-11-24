package dbrepo

import (
	"database/sql"

	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.Repository {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
