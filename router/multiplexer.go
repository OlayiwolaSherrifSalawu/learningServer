package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Application struct {
	Store    *Store
	FileName string
}
type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

type Store map[int]Task

func (app *Application) HandleTask(w http.ResponseWriter, r *http.Request) {
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
	task, ok := (*app.Store)[id]
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
func (app *Application) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only post method allowed"))
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 104567)
	count := 0
	for _, val := range *app.Store {
		if val.Id >= count {
			count = val.Id
		}
	}
	count += 1
	newTask := (*app.Store)[count]

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
	(*app.Store)[count] = newTask
	app.Store.WriteJson(app.FileName)
}

func (app *Application) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", "PUT")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 104567)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not found"))
		fmt.Println(err)
		return
	}
	taskToUpdate, ok := (*app.Store)[id]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID not found "))
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&taskToUpdate); err != nil {
		http.Error(w, "error while decoding", http.StatusBadRequest)
		return
	}
	taskToUpdate.Id = id
	json.NewEncoder(w).Encode(taskToUpdate)
	(*app.Store)[id] = taskToUpdate

	app.Store.WriteJson(app.FileName)
}

func (app *Application) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", "Only DELETE")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only Delete Method Allowed"))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(" id not allowed "))
		return
	}
	_, ok := (*app.Store)[id]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(" id not allowed cant find id to delete"))
		return
	}
	delete((*app.Store), id)
	app.Store.WriteJson(app.FileName)
}
