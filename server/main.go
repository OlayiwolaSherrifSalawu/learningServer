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
	Id          int    `json:"id"`
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
	newStore := make(Store)
	err := json.Unmarshal(slices, toMarshal)
	if err != nil {
		return err
	}
	for _, val := range *toMarshal {
		(newStore)[val.Id] = val
	}
	(*s) = newStore
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
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	task, ok := (*app.store)[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	byts, err := json.Marshal(task)
	if err != nil {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(501)
		w.Write([]byte("error reading files "))
		return
	}
	w.Write(byts)
}
func (app *Application) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "only post methods allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only post method allowed"))
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 104567)
	count := 0
	for _, val := range *app.store {
		if val.Id >= count {
			count = val.Id
		}
	}
	count += 1
	newTask := (*app.store)[count]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newTask); err != nil {
		http.Error(w, "error decoding the message", http.StatusBadRequest)
		return
	}
	newTask.Id = count

	err := json.NewEncoder(w).Encode(newTask)
	if err != nil {
		http.Error(w, "error encoding the message", http.StatusBadRequest)
		return
	}
	(*app.store)[count] = newTask
	app.store.writeJson(app.fileName)
}

func (s Store) writeJson(fileName string) error {
	var theStore []Task
	for _, val := range s {
		theStore = append(theStore, val)
	}

	byts, err := json.Marshal(theStore)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, byts, 0644)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	var temp []Task
	theStore := &Store{}
	fileByte, err := theStore.readJson("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	theStore.unmarsha(fileByte, &temp)
	app := &Application{
		store:    theStore,
		fileName: "result.json",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.handleTask)
	mux.HandleFunc("POST /", app.createTask)
	log.Println("server started at port 8080")
	errs := http.ListenAndServe(":8080", mux)
	log.Fatal(errs)
}
