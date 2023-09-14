-- name: CreateItem :one
INSERT INTO items(user_id, amount, kind, happened_at, tag_ids)
  VALUES ($1, $2, $3, $4, $5)
RETURNING
  *;

-- name: ListItem :many
SELECT
  *
FROM
  items
WHERE
  user_id = $1
  AND happened_at >= sqlc.arg(happened_at_begin)
  AND happened_at <= sqlc.arg(happened_at_end)
ORDER BY
  happened_at DESC offset $2
LIMIT $3;

-- name: CountItem :one
SELECT
  count(*)
FROM
  items
WHERE
  user_id = $1;

