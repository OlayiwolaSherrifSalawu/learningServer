package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
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
	data := &templatesData{Snippet: snippet}
	app.render(w, r, "show.page.tmpl", data)
	
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

func newCacheTemplate(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partail.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
