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
	store    *Store
	fileName string
}
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

type Store map[int]Task

func (s *Store) readJson(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (s *Store) unmarsha(slices []byte, toMarshal *[]Task) error {

	err := json.Unmarshal(slices, toMarshal)
	if err != nil {
		return err
	}
	for i, val := range *toMarshal {
		(*s)[i] = val
	}
	return nil
}
func (app *Application) handleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	task := (*app.store)[id]
	byts, err := json.Marshal(task)
	if err != nil {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(501)
		w.Write([]byte("error reading files "))
		return
	}
	w.Write(byts)
}

func main() {
	var temp *[]Task
	theStore := &Store{}
	fileByte, err := theStore.readJson("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	theStore.unmarsha(fileByte, temp)
	app := &Application{
		store: theStore,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleTask)
	log.Println("server started at port 8080")
	errs := http.ListenAndServe(":8080", mux)
	log.Fatal(errs)
}
