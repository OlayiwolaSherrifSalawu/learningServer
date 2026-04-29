package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func (app *Application) routes(cfg *Config) (http.Handler, *Config) {
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Address")
	app.ErrorLoger = log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	app.InfoLogger = log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.CreateSnippet)
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", fileServer))
	return app.recoverOnPanic(app.logRequest(secureHeaders(mux))), cfg
}
