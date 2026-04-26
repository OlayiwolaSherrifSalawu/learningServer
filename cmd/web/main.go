package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"alexedwards.net/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr       string
	StaticAddr string
}

type Application struct {
	ErrorLoger    *log.Logger
	InfoLogger    *log.Logger
	snippet       *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	cfg := new(Config)

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MY SQL DSN")
	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	tsCache, err := newCacheTemplate("ui/html")
	if err != nil {
		fmt.Println(err)
		return
	}

	app := &Application{
		snippet:       &mysql.SnippetModel{DB: db},
		templateCache: tsCache,
	}
	mux, cfg := app.routes(cfg)
	serve := &http.Server{
		Addr:     cfg.Addr,
		Handler:  mux,
		ErrorLog: app.ErrorLoger,
	}
	app.InfoLogger.Printf("Starting server on %s", cfg.Addr)
	err = serve.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		serve.ErrorLog.Println("server closed")
		return
	}
	serve.ErrorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
