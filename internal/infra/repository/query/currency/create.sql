-- name: Create :one
INSERT INTO currency
    (name, code, usd_exchange_rate, created_by, "number", country, country_code, search_url)
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;