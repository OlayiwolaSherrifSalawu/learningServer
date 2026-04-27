package main

import (
	"alexedwards.net/snippetbox/pkg/models"
)

type templatesData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}
