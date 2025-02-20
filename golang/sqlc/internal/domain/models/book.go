package models

import (
	"time"
)

type Book struct {
	ID            int64
	Title         string
	AuthorID      int64
	PublishedDate time.Time
}
