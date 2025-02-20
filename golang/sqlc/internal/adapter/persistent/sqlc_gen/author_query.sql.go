// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: author_query.sql

package sqlc_gen

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id, name, bio, created_at, updated_at
`

type CreateAuthorParams struct {
	Name string
	Bio  pgtype.Text
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAuthor, id)
	return err
}

const filterAuthors = `-- name: FilterAuthors :many
SELECT id, name, bio, created_at, updated_at
FROM authors
WHERE (
        COALESCE($3, '') = ''
        OR name ILIKE '%' || $3 || '%'
    ) -- search by name
    AND (
        COALESCE($4, '') = ''
        OR bio ILIKE '%' || $4 || '%'
    ) -- search by bio
ORDER BY name
LIMIT $1 OFFSET $2
`

type FilterAuthorsParams struct {
	Limit  int32
	Offset int32
	Name   interface{}
	Bio    interface{}
}

func (q *Queries) FilterAuthors(ctx context.Context, arg FilterAuthorsParams) ([]Author, error) {
	rows, err := q.db.Query(ctx, filterAuthors,
		arg.Limit,
		arg.Offset,
		arg.Name,
		arg.Bio,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio, created_at, updated_at FROM authors WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBooksByAuthorID = `-- name: GetBooksByAuthorID :many
SELECT id, title, author_id, published_date FROM books
WHERE id = $1
`

func (q *Queries) GetBooksByAuthorID(ctx context.Context, id int64) ([]Book, error) {
	rows, err := q.db.Query(ctx, getBooksByAuthorID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.AuthorID,
			&i.PublishedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio, created_at, updated_at FROM authors ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
SET
    name = $2,
    bio = $3,
    updated_at = NOW()
WHERE
    id = $1
`

type UpdateAuthorParams struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.Exec(ctx, updateAuthor, arg.ID, arg.Name, arg.Bio)
	return err
}
