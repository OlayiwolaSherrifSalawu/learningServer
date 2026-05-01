package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"alexedwards.net/snippetbox/pkg/models"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	late, err := app.snippet.Latest()
	data := &templatesData{Snippets: late}
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", data)
}
func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	data := &templatesData{Snippet: snippet}
	app.render(w, r, "show.page.tmpl", data)

}
func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
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
func (app *Application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
