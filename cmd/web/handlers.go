package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"alexedwards.net/snippetbox/pkg/forms"
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
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")

	if title == "" {
		http.Error(w, "title cannot be empty!", http.StatusBadRequest)
		return
	}

	formss := forms.NewForm(r.PostForm)

	formss.Required("title", "content", "expires")
	formss.PermittedValues("expires", "1", "7", "365")
	formss.MaxLength("title", 100)

	// Initialize a map to hold any validation errors.
	if !formss.Valid() {
		app.render(w, r, "create.page.tmpl", &templatesData{
			Forms: formss,
		})
		return
	}
	id, err := app.snippet.Insert(formss.Get("title"), formss.Get("content"), formss.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
	// w.Write([]byte("Create a new snippet..."))
}
func (app *Application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templatesData{
		Forms: forms.NewForm(nil),
	})
}
