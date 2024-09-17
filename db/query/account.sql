-- name: GetAllAccount :many
SELECT * FROM "account"
LIMIT $1;