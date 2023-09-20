// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: item.sql

package queries

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const countItem = `-- name: CountItem :one
SELECT
  count(*)
FROM
  items
WHERE
  user_id = $1
`

func (q *Queries) CountItem(ctx context.Context, userID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countItem, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createItem = `-- name: CreateItem :one
INSERT INTO items(user_id, amount, kind, happened_at, tag_ids)
  VALUES ($1, $2, $3, $4, $5)
RETURNING
  id, user_id, amount, tag_ids, kind, happened_at, created_at, updated_at
`

type CreateItemParams struct {
	UserID     int32     `json:"userId"`
	Amount     int32     `json:"amount"`
	Kind       Kind      `json:"kind"`
	HappenedAt time.Time `json:"happenedAt"`
	TagIds     []int32   `json:"tagIds"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.UserID,
		arg.Amount,
		arg.Kind,
		arg.HappenedAt,
		pq.Array(arg.TagIds),
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		pq.Array(&i.TagIds),
		&i.Kind,
		&i.HappenedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
WHERE user_id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteItem, userID)
	return err
}

const itemsBalance = `-- name: ItemsBalance :many
SELECT
  amount
  -- SUM(amount)
FROM
  items
WHERE
  kind = $1
  AND user_id = $2
  AND happened_at >= $3
  AND happened_at < $4
`

type ItemsBalanceParams struct {
	Kind            Kind      `json:"kind"`
	UserID          int32     `json:"userId"`
	HappenedAtBegin time.Time `json:"happenedAtBegin"`
	HappenedAtEnd   time.Time `json:"happenedAtEnd"`
}

func (q *Queries) ItemsBalance(ctx context.Context, arg ItemsBalanceParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, itemsBalance,
		arg.Kind,
		arg.UserID,
		arg.HappenedAtBegin,
		arg.HappenedAtEnd,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var amount int32
		if err := rows.Scan(&amount); err != nil {
			return nil, err
		}
		items = append(items, amount)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listItem = `-- name: ListItem :many
SELECT
  id, user_id, amount, tag_ids, kind, happened_at, created_at, updated_at
FROM
  items
WHERE
  user_id = $1
  AND happened_at >= $4
  AND happened_at < $5
ORDER BY
  happened_at DESC offset $2
LIMIT $3
`

type ListItemParams struct {
	UserID          int32     `json:"userId"`
	Offset          int32     `json:"offset"`
	Limit           int32     `json:"limit"`
	HappenedAtBegin time.Time `json:"happenedAtBegin"`
	HappenedAtEnd   time.Time `json:"happenedAtEnd"`
}

func (q *Queries) ListItem(ctx context.Context, arg ListItemParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItem,
		arg.UserID,
		arg.Offset,
		arg.Limit,
		arg.HappenedAtBegin,
		arg.HappenedAtEnd,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			pq.Array(&i.TagIds),
			&i.Kind,
			&i.HappenedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listItemsByHappenedAtAndKind = `-- name: ListItemsByHappenedAtAndKind :many
SELECT
  id, user_id, amount, tag_ids, kind, happened_at, created_at, updated_at
FROM
  items
WHERE
  happened_at >= $2
  AND happened_at < $3
  AND kind = $1
  AND user_id = $4
ORDER BY
  happened_at DESC
`

type ListItemsByHappenedAtAndKindParams struct {
	Kind            Kind      `json:"kind"`
	HappenedAtBegin time.Time `json:"happenedAtBegin"`
	HappenedAtEnd   time.Time `json:"happenedAtEnd"`
	UserID          int32     `json:"userId"`
}

func (q *Queries) ListItemsByHappenedAtAndKind(ctx context.Context, arg ListItemsByHappenedAtAndKindParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItemsByHappenedAtAndKind,
		arg.Kind,
		arg.HappenedAtBegin,
		arg.HappenedAtEnd,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			pq.Array(&i.TagIds),
			&i.Kind,
			&i.HappenedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
