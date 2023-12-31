// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES($1, $2, $3)
RETURNING id, created_at, updated_at, name, api_key
`

type CreateUserParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.CreatedAt, arg.UpdatedAt, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, created_at, updated_at, name, api_key FROM users WHERE api_key = $1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
