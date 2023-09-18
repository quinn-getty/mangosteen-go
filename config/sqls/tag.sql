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

-- name: UpdateTag :one
UPDATE
  tags
SET
  name = CASE WHEN @name::varchar = '' THEN
    name
  ELSE
    @name
  END,
  sign = CASE WHEN @sign::varchar = '' THEN
    sign
  ELSE
    @sign
  END,
  updated_at = CASE WHEN @updated_at::timestamp = '' THEN
    updated_at
  ELSE
    @updated_at
  END,
  updated_at = now()
WHERE
  id = $1
RETURNING
  *;

-- name: DeleteTag :one
UPDATE
  tags
SET
  deleted_at = now(),
  updated_at = now()
WHERE
  id = $1
RETURNING
  *;

-- name: DeleteUserAllTag :exec
DELETE FROM tags
WHERE user_id = $1;

-- name: FindTag :one
SELECT
  *
FROM
  tags
WHERE
  id = $1
  AND user_id = $2
  AND deleted_at IS NULL;

