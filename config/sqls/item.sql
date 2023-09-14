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
  AND happened_at >= $2
  AND happened_at <= $3
ORDER BY
  happened_at DESC offset $4
LIMIT $5;

-- name: CountItem :one
SELECT
  count(*)
FROM
  items
WHERE
  user_id = $1;

