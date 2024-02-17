-- name: CreatePost :one
INSERT INTO posts (title, description, published_at, url, feed_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetPostsByFeedFollow :many
SELECT * FROM posts WHERE feed_id = ANY($1::int[]) ORDER BY published_at DESC LIMIT $2;