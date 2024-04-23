-- name: CreateAccount :one
INSERT INTO accounts (
    name, email, phone, password
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetListAccount :many
SELECT id, name, email, phone, password_changed_at, created_at, updated_at, is_email_verified
FROM accounts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetDetailAccount :one
SELECT id, name, email, phone, password_changed_at, created_at, updated_at, is_email_verified
FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
SET 
    name = $2,
    email = $3,
    phone = $4,
    password = $5,
    password_changed_at = $6,
    updated_at = CURRENT_TIMESTAMP,
    is_email_verified = $7
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM accounts
WHERE id = $1
RETURNING *;