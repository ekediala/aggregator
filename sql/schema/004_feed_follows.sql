-- +goose Up
CREATE TABLE feed_follows (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id SERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  feed_id BIGSERIAL NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows