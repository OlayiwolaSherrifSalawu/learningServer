package models

import (
	"errors"
	"time"
)

var ErrNoRecods = errors.New("No matching records found")

// a table struct that represent and correspond with our table
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
