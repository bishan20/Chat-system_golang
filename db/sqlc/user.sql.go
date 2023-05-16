// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name
) VALUES (
  $1
) RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
