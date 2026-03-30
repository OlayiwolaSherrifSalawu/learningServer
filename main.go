package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Store interface {
	Save()
	Delete()
}

type Task struct {
	ID         int    `json:"id"`
	TITLE      string `json:"title"`
	IsComplete bool   `json:"isComplete"`
}
type Logger struct{}

func (l Logger) Save() {

}
func (l Logger) Delete() {

}
func (t *Task) Save() {
	t.IsComplete = true
}
func (t Task) Delete() {

}
func ProcessInterface(s Store) {
	s.Save()
}
func main() {
	task1 := &Task{
		ID:         38,
		TITLE:      "read through the length of string ",
		IsComplete: false,
	}
	ProcessInterface(task1)
	js, err := json.Marshal(task1)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = os.WriteFile("user.json", js, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
