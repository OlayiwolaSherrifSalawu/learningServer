package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Application struct {
	store *Store
}
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type Store map[int]Task

func (app *Application) handleTask(w http.ResponseWriter, r *http.Request) {
task:= app.store[]
}

func (s *Store) readJson(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, s)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) getJson(id int, fileName string) ([]byte, error) {
	err := s.readJson(fileName)
	if err != nil {
		return nil, err
	}

	val, ok := (*s)[id]
	if !ok {
		return nil, fmt.Errorf("task cant be found ")
	}
	byts, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return byts, nil
}
func task(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// data, err := json.Marshal(file
	// w.Write(file)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", task)
	log.Println("started server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
