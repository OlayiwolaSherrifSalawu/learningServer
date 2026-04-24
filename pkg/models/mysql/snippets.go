package mysql

import (
	"database/sql"
	"errors"

	"alexedwards.net/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

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
	// stmt := `SELECT id, title, content, created, expires FROM snippets
	// WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// row := s.DB.QueryRow(stmt, ID)

	snipps := &models.Snippet{}
	err := s.DB.QueryRow("SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?", ID).Scan(&snipps.ID, &snipps.Title, &snipps.Content, &snipps.Created, &snipps.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecods
		} else {
			return nil, err
		}
	}
	return snipps, nil
}
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		ss := &models.Snippet{}
		err := rows.Scan(&ss.ID, &ss.Title, &ss.Content, &ss.Created, &ss.Created, &ss.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, ss)
	}

	return snippets, nil
}
