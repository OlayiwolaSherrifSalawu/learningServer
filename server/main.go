package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type Store map[int]Task

func (s *Store) readJson(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, s)
	if err != nil {
		fmt.Println("afang")
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
	var s Store
	err := s.readJson("result.json")
	if err != nil {
		w.WriteHeader(500)
		w.Header().Add("Allow", "Error reading file")
		w.Write([]byte("server error while reading file "))
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(501)
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

	// data, err := json.Marshal(file)
	file, err := s.getJson(id, "result.json")
	if err != nil {
		w.Write([]byte("something went wrong"))
		return
	}

	w.Write(file)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", task)
	log.Println("started server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
