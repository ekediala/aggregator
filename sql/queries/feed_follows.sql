-- name: FollowFeed :one
INSERT INTO feed_follows(user_id, feed_id)
VALUES($1, $2)
RETURNING *;

-- name: UnfollowFeed :execrows
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2;

-- name: GetUserFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;