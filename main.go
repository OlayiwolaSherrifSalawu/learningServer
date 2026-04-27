package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hell0 world!, %s", r.URL.Path[1:])
}
func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handlerFunc)
	// err := http.ListenAndServe(":8080", mux)
	// log.Fatal(err)

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
