-- name: Update :exec
UPDATE
    currency
SET
    id = $2,
    name = $3,
    code = $4,
    number = $5,
    country = $6,
    country_code = $7,
    search_url = $8,
    usd_exchange_rate = $9,
    fic = $10,
    updated_by = $11,
    updated_at = $12
WHERE
    id = $1;