-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, api_key)
VALUES (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPiKey :one
SELECT * from users WHERE api_key = $1;
