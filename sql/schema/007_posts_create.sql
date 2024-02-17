-- +goose Up
CREATE TABLE posts (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  url VARCHAR(100) NOT NULL UNIQUE,
  feed_id BIGSERIAL REFERENCES feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts;