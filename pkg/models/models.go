package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
