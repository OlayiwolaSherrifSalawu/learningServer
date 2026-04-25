package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"alexedwards.net/snippetbox/pkg/models"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// files := []string{
	// 	"ui/html/home.page.tmpl",
	// 	"ui/html/base.layout.tmpl",
	// 	"ui/html/footer.partail.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, nil)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	late, err := app.snippet.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, val := range late {
		fmt.Fprintf(w, "%v", val)
	}
}
func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippet.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecods) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}
	files := []string{
		"ui/html/show.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partail.tmpl",
	}
	temples, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = temples.Execute(w, snippet)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Ola test"
	content := "Afang at night "
	expires := "3"

	id, err := app.snippet.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	// w.Write([]byte("Create a new snippet..."))
}
