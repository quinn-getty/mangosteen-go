-- name: CreateValidationCode :one
INSERT INTO validation_codes(email, code)
  VALUES ($1, $2)
RETURNING
  *;

