-- name: CreateBooks :copyfrom
INSERT INTO books (title, author_id, published_date) VALUES ($1, $2, $3);