-- name: CreateMember :one
INSERT INTO members (full_name, email, phone)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMemberByID :one
SELECT * FROM members WHERE id = $1;

-- name: ListMembers :many
SELECT * FROM members ORDER BY created_at DESC LIMIT $1 OFFSET $2;