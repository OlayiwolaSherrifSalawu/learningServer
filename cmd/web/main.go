package main

import (
	"log"
	"net/http"
)

type Config struct {
	Addr       string
	StaticAddr string
}

type Application struct {
	ErrorLoger *log.Logger
	InfoLogger *log.Logger
}

func main() {
	cfg := new(Config)
	app := new(Application)
	mux, cfg := app.routes(cfg)
	serve := &http.Server{
		Addr:     cfg.Addr,
		Handler:  mux,
		ErrorLog: app.ErrorLoger,
	}
	app.InfoLogger.Printf("Starting server on %s", cfg.Addr)
	err := serve.ListenAndServe()
	serve.ErrorLog.Fatal(err)

}
