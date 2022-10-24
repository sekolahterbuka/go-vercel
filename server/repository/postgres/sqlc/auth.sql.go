// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: auth.sql

package pg

import (
	"context"

	"github.com/google/uuid"
)

const login = `-- name: Login :one

SELECT id, username, password, email FROM auth WHERE id = $1 LIMIT 1
`

func (q *Queries) Login(ctx context.Context, id uuid.UUID) (Auth, error) {
	row := q.queryRow(ctx, q.loginStmt, login, id)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
	)
	return i, err
}

const register = `-- name: Register :one

INSERT INTO
    auth (username, email, password)
VALUES ($1, $2, $3) RETURNING id, username, password, email
`

type RegisterParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) Register(ctx context.Context, arg RegisterParams) (Auth, error) {
	row := q.queryRow(ctx, q.registerStmt, register, arg.Username, arg.Email, arg.Password)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
	)
	return i, err
}
