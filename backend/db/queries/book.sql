-- name: CreateBook :one
INSERT INTO books (title, isbn, openlibrary_id, cover_url, category_id, stock, available)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetBookByID :one
SELECT * FROM books WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: UpdateBook :one
UPDATE books
SET title = $2, isbn = $3, cover_url = $4, category_id = $5, stock = $6, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;