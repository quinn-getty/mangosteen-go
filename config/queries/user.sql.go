// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package queries

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(email)
  VALUES ($1)
RETURNING
  id, email, phone, address, created_at, updated_at
`

func (q *Queries) CreateUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT
  id, email, phone, address, created_at, updated_at
FROM
  users
WHERE
  email = $1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT
  id, email, phone, address, created_at, updated_at
FROM
  users
WHERE
  id = $1
`

func (q *Queries) FindUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT
  id, email, phone, address, created_at, updated_at
FROM
  users
ORDER BY
  id OFFSET $1
LIMIT $2
`

type ListUsersParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Phone,
			&i.Address,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE
  users
SET
  email = $2,
  phone = $3,
  address = $4
WHERE
  id = $1
`

type UpdateUserParams struct {
	ID      int32
	Email   string
	Phone   string
	Address string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Phone,
		arg.Address,
	)
	return err
}
