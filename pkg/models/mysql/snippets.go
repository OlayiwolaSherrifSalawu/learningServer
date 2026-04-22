package mysql

import (
	"database/sql"

	"alexedwards.net/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title, content, created, expires string) (int, error) {
	return 0, nil
}

func (s *SnippetModel) Get(ID int) (*models.Snippet, error) {
	return nil, nil
}
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
