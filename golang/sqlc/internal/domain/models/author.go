package models

import "time"

type Author struct {
	ID        int64
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []Book
}

type AuthorFilter struct {
	Name string
	Bio  string
}
