package main

import (
	"alexedwards.net/snippetbox/pkg/models"
)

type templatesData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
