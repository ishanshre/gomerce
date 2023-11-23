package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ishanshre/gomerce/internals/config"
)

var app config.AppConfig

func main() {
	flag.IntVar(&app.Port, "port", 8000, "Port for server to listen to")
	flag.Parse()

	app.Addr = fmt.Sprintf(":%d", app.Port)
	srv := http.Server{
		Addr:    app.Addr,
		Handler: nil,
	}
	log.Printf("Starting server at port %d", app.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
