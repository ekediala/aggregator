package main

import (
	"database/sql"
	"time"
)

type ReturnedUser struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type ReturnedFeed struct {
	ID        int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    int32 `json:"user_id"`
}

type ReturnedFeedFollow struct {
	ID        int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int32 `json:"user_id"`
	FeedID    int64 `json:"feed_id"`
}

type ReturnedPost struct {
	ID          int64 `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string `json:"title"`
	Description sql.NullString `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string `json:"url"`
	FeedID      int64 `json:"feed_id"`
}
