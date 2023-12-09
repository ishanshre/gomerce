package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ishanshre/gomerce/internals/config"
	"github.com/ishanshre/gomerce/internals/connection"
	"github.com/ishanshre/gomerce/internals/handler"
	"github.com/ishanshre/gomerce/internals/middleware"
	"github.com/ishanshre/gomerce/internals/repository/dbrepo"
	"github.com/ishanshre/gomerce/internals/router"
	"github.com/joho/godotenv"
)

var app config.AppConfig

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error in loading environment files: %s\n", err.Error())
	}
	flag.IntVar(&app.Port, "port", 8000, "Port for server to listen to")
	flag.Parse()

	app.DbString = os.Getenv("driver")
	app.Dsn = os.Getenv("DB_URL")
	// app.RedisHost = os.Getenv("redis")
	app.Addr = fmt.Sprintf(":%d", app.Port)

	handler, middleware, connection := run(&app, context.Background())

	defer connection.CloseDb()

	srv := http.Server{
		Addr:    app.Addr,
		Handler: router.Router(&app, handler, middleware),
	}
	app.InfoLog.Printf("Starting server at port %d", app.Port)
	if err := srv.ListenAndServe(); err != nil {
		app.ErrorLog.Fatalf("error: %s", err.Error())
	}
}

func run(app *config.AppConfig, ctx context.Context) (handler.Handler, middleware.Middleware, connection.Connection) {
	app.InProduction = false
	app.UseCache = false
	app.UseRedis = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	conn := connection.NewConnection(app.DbString, app.Dsn, ctx)

	repo := dbrepo.NewPostgresRepo(conn, app, ctx)

	handler := handler.NewHandler(app, repo, conn, ctx)

	middleware := middleware.NewMiddleware(app, ctx)

	return handler, middleware, conn
}
