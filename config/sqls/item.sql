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
  AND happened_at < sqlc.arg(happened_at_end)
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

-- name: ItemsBalance :many
SELECT
  amount
  -- SUM(amount)
FROM
  items
WHERE
  kind = $1
  AND user_id = $2
  AND happened_at >= sqlc.arg(happened_at_begin)
  AND happened_at < sqlc.arg(happened_at_end);

-- name: DeleteItem :exec
DELETE FROM items
WHERE user_id = $1;

-- name: ListItemsByHappenedAtAndKind :many
SELECT
  *
FROM
  items
WHERE
  happened_at >= @happened_at_begin
  AND happened_at < @happened_at_end
  AND kind = $1
  AND user_id = @user_id
ORDER BY
  happened_at DESC;

