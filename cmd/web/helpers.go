package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n %s", err.Error(), debug.Stack())
	app.ErrorLoger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, name string, td *templatesData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("No templateCache"))
		return
	}

	buff := new(bytes.Buffer)

	err := ts.Execute(buff, app.addDefaultData(td,r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	buff.WriteTo(w)
}

func (app *Application) addDefaultData(td *templatesData, r *http.Request) *templatesData {
	if td == nil {
		td = &templatesData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}
