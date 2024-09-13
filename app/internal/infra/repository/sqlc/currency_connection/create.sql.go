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
    (name, code, usd_exchange_rate, created_by, "number", country, country_code, search_url, response_path_to_rate, updated_by)
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING id, name, code, number, country, country_code, usd_exchange_rate, search_url, response_path_to_rate, fic, created_by, updated_by, created_at, updated_at
`

type CreateParams struct {
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	UsdExchangeRate    float64 `json:"usd_exchange_rate"`
	CreatedBy          int     `json:"created_by"`
	Number             *int    `json:"number"`
	Country            *string `json:"country"`
	CountryCode        *string `json:"country_code"`
	SearchUrl          *string `json:"search_url"`
	ResponsePathToRate *string `json:"response_path_to_rate"`
	UpdatedBy          *int    `json:"updated_by"`
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
		arg.ResponsePathToRate,
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
		&i.UsdExchangeRate,
		&i.SearchUrl,
		&i.ResponsePathToRate,
		&i.Fic,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
