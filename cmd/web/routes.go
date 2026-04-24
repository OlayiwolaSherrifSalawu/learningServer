package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func (app *Application) routes(cfg *Config) (*http.ServeMux, *Config) {
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Address")
	flag.StringVar(&cfg.StaticAddr, "staticAddr", "ui/static/", "Static Files")

	flag.Parse()
	app.ErrorLoger = log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	app.InfoLogger = log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	fileServer := http.FileServer(http.Dir("/static/"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.CreateSnippet)
	mux.Handle("/static/", http.StripPrefix("static/", fileServer))
	return mux, cfg
}
