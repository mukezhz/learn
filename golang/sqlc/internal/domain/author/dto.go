package author

import (
	"time"

	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/models"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
)

type CreateAuthorParams struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio" binding:"required"`
}

func (p CreateAuthorParams) ToModel() models.Author {
	return models.Author{
		Name: p.Name,
		Bio:  p.Bio,
	}
}

type AuthorResponse struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Bio       string         `json:"bio"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Books     []BookResponse `json:"books"`
}

type FilterAuthorsParams struct {
	framework.PaginatedRequest
	Name string `form:"name"`
	Bio  string `form:"bio"`
}

func (p FilterAuthorsParams) ToModel() models.AuthorFilter {
	return models.AuthorFilter{
		Name: p.Name,
		Bio:  p.Bio,
	}
}

type BookResponse struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	PublishedDate time.Time `json:"published_date"`
}

func (BookResponse) FromModel(b models.Book) BookResponse {
	return BookResponse{
		ID:            b.ID,
		Title:         b.Title,
		PublishedDate: b.PublishedDate,
	}
}

type CreateAuthorWithBooksParams struct {
	CreateAuthorParams
	Books []CreateBookParams `json:"books" binding:"required"`
}

type CreateBookParams struct {
	Title         string    `json:"title" binding:"required"`
	PublishedDate time.Time `json:"published_date" binding:"required"`
}

func (p CreateBookParams) ToModel() models.Book {
	return models.Book{
		Title:         p.Title,
		PublishedDate: p.PublishedDate,
	}
}
