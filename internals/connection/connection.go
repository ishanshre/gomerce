package connection

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Connection interface{}

type connection struct {
	SQL *sql.DB
	ctx context.Context
}

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxLifeDBTime = 5 * time.Minute
)

func NewConnection(dbString, dsn string, ctx context.Context) Connection {
	db, err := newDatabase(dbString, dsn)
	if err != nil {
		log.Fatal("DB error: ", err)
	}
	return &connection{
		SQL: db,
		ctx: ctx,
	}
}

func newDatabase(dbString, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbString, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxLifeDBTime)
	return db, err
}
