package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *Application) routes(cfg *Config) (http.Handler, *Config) {
	standardMid := alice.New(app.recoverOnPanic, app.logRequest, secureHeaders)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Address")
	app.ErrorLoger = log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	app.InfoLogger = log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	fileServer := http.FileServer(http.Dir("ui/static/"))

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", fileServer))
	return standardMid.Then(mux), cfg
}
