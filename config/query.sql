-- name: CreateUser :one
INSERT INTO users(email)
  VALUES ($1)
RETURNING
  *;

-- name: UpdateUser :exec
UPDATE
  users
SET
  email = $2,
  phone = $3,
  address = $4
WHERE
  id = $1;

