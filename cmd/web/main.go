package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
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
	ErrorLoger *log.Logger
	InfoLogger *log.Logger
	snippet    *mysql.SnippetModel
}

func main() {
	cfg := new(Config)
	app := new(Application)
	mux, cfg := app.routes(cfg)
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MY SQL DSN")
	flag.Parse()
	db, err := openDB(*dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	app.snippet = &mysql.SnippetModel{DB: db}

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
