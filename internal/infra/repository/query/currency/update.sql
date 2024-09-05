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
    real_time_value = $9,
    usd_exchange_rate = $10,
    fic = $11,
    updated_by = $12,
    updated_at = $13
WHERE
    id = $1;