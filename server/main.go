package main

import (
	"fmt"
	"log"
	"net/http"

	"alexedwards.net/snippetbox/router"
)

func main() {
	var temp []router.Task
	theStore := &router.Store{}
	fileByte, err := theStore.ReadJson("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	theStore.Unmarsha(fileByte, &temp)
	app := &router.Application{
		Store:    theStore,
		FileName: "result.json",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.HandleTask)
	mux.HandleFunc("POST /", app.CreateTask)
	mux.HandleFunc("PUT /", app.UpdateTask)
	mux.HandleFunc("DELETE /", app.DeleteTask)
	log.Println("server started at port 8080")
	errs := http.ListenAndServe(":8080", mux)
	log.Fatal(errs)
}
