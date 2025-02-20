-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
SET
    name = $2,
    bio = $3,
    updated_at = NOW()
WHERE
    id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = $1;

-- name: FilterAuthors :many
SELECT *
FROM authors
WHERE (
        @name::TEXT IS NULL
        OR name ILIKE '%' || @name || '%'
    ) -- search by name
    AND (
        @bio::TEXT IS NULL
        OR bio ILIKE '%' || @bio || '%'
    ) -- search by bio
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: GetBooksByAuthorID :many
SELECT * FROM books
WHERE id = $1;  