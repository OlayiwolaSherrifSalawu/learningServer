package mysql

import (
	"database/sql"

	"alexedwards.net/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title, content,  expires string) (int, error) {
	stmt := `INSERT INTO snippets (title,content,created,expires) VALUES(?,?, UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL?DAY))`

	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *SnippetModel) Get(ID int) (*models.Snippet, error) {
	return nil, nil
}
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
