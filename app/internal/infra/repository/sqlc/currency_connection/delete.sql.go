// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: delete.sql

package currency_connection

import (
	"context"
)

const delete = `-- name: Delete :exec
DELETE FROM currency WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}
