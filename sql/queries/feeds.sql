-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, name, url, user_id)
VALUES (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, $2, $3)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: UpdateLastFetchedAt :one
UPDATE feeds SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * from feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;