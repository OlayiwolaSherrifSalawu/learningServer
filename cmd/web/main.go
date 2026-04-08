package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("ui/static/"))

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Starting server on :4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
