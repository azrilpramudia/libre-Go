-- name: CreateMember :one
INSERT INTO members (full_name, email, phone, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMemberByID :one
SELECT * FROM members WHERE id = $1;

-- name: GetMemberByEmail :one
SELECT * FROM members WHERE email = $1;

-- name: ListMembers :many
SELECT * FROM members ORDER BY created_at DESC LIMIT $1 OFFSET $2;