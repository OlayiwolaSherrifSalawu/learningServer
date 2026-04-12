package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr       string
	StaticAddr string
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticAddr, "static", "ui/static/", "Static files")

	infoLogger := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errors := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	flag.Parse()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.StaticAddr))

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	serve := &http.Server{
		Addr:     cfg.Addr,
		Handler:  mux,
		ErrorLog: errors,
	}
	infoLogger.Printf("Starting server on %s", cfg.Addr)
	err := serve.ListenAndServe()
	serve.ErrorLog.Fatal(err)

}
