-- name: GetAuthorByName :one
SELECT * FROM authors WHERE name = $1;

-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES ($1) RETURNING *;