-- name: CreateTag :one
INSERT INTO tags(user_id, name, sign)
  VALUES ($1, $2, $3)
RETURNING
  *;

-- name: ListTag :many
SELECT
  *
FROM
  tags
WHERE
  user_id = $1
  AND deleted_at IS NULL
ORDER BY
  created_at DESC;

