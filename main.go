package main

import (
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello world"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunc)

	log.Print("started server at :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
