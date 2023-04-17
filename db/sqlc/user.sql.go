// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    hashedpassword
) VALUES (
             $1, $2, $3
         ) RETURNING id, name, email, hashedpassword, created_at
`

type CreateUserParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Hashedpassword string `json:"hashedpassword"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Hashedpassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, hashedpassword, created_at FROM users
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
	)
	return i, err
}