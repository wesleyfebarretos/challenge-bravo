-- name: Update :exec
UPDATE
    currency
SET
    name = $2,
    code = $3,
    number = $4,
    country = $5,
    country_code = $6,
    search_url = $7,
    usd_exchange_rate = $8,
    fic = $9,
    updated_by = $10,
    updated_at = $11
WHERE
    id = $1;