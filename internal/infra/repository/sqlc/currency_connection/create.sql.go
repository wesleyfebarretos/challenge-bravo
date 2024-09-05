// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: create.sql

package currency_connection

import (
	"context"
)

const create = `-- name: Create :one
INSERT INTO currency
    (name, code, usd_exchange_rate, created_by, "number", country, country_code, search_url, updated_by)
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING id, name, code, number, country, country_code, search_url, usd_exchange_rate, fic, created_by, updated_by, created_at, updated_at
`

type CreateParams struct {
	Name            string  `json:"name"`
	Code            string  `json:"code"`
	UsdExchangeRate float64 `json:"usd_exchange_rate"`
	CreatedBy       int     `json:"created_by"`
	Number          *int    `json:"number"`
	Country         *string `json:"country"`
	CountryCode     *string `json:"country_code"`
	SearchUrl       *string `json:"search_url"`
	UpdatedBy       *int    `json:"updated_by"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Currency, error) {
	row := q.db.QueryRow(ctx, create,
		arg.Name,
		arg.Code,
		arg.UsdExchangeRate,
		arg.CreatedBy,
		arg.Number,
		arg.Country,
		arg.CountryCode,
		arg.SearchUrl,
		arg.UpdatedBy,
	)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Number,
		&i.Country,
		&i.CountryCode,
		&i.SearchUrl,
		&i.UsdExchangeRate,
		&i.Fic,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
