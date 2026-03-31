package main

import (
	"log"
	"net/http"
	"os"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type Store map[string]Task

func (s Store) getJson(id string) []byte {

	return nil
}
func task(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("result.json")
	if err != nil {
		w.WriteHeader(500)
		w.Header().Add("Allow", "Error reading file")
		w.Write([]byte("server error while reading file "))
		return
	}
	if r.Method != http.MethodGet {
		w.Write([]byte("method not allowed"))
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// data, err := json.Marshal(file)

	w.Write(file)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", task)
	log.Println("started server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
