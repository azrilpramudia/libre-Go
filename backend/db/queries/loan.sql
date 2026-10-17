-- name: CreateLoan :one
INSERT INTO loans (book_id, member_id, due_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLoanByID :one
SELECT * FROM loans WHERE id = $1;

-- name: ReturnLoan :one
UPDATE loans
SET returned_at = now(), status = 'returned', fine_amount = $2
WHERE id = $1 AND status = 'borrowed'
RETURNING *;

-- name: ListLoansByMember :many
SELECT * FROM loans WHERE member_id = $1 ORDER BY borrowed_at DESC;