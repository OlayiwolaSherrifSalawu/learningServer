package main

import (
	"encoding/json"
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

func task(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("result.json")
	if err != nil {
		w.WriteHeader(501)
		w.Header().Add("Error", "Error reading file")
		w.Write([]byte("server error while reading file "))
	}
	if r.Method != http.MethodGet {
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	data, err := json.Marshal(file)
	if err != nil {
		w.WriteHeader(501)
		w.Header().Add("Error", "Error reading file")
		w.Write([]byte("server error while reading file "))
	}
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", task)
	log.Println("started server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
