package main

import (
	"errors"
	"net/http"

	"github.com/ekediala/aggregator/internal/database"
)

func DatabaseUserToReturnedUser(user database.User) ReturnedUser {
	return ReturnedUser{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func DatabaseFeedFollowToResponseFeedFollow(feedFollow database.FeedFollow) ReturnedFeedFollow {
	return ReturnedFeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func DatabaseFeedToReturnedFeed(feed database.Feed) ReturnedFeed {
	return ReturnedFeed{
		ID:        feed.ID,
		Name:      feed.Name,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		Url:       feed.Url,
	}
}

func DatabasePostToReturnedPost(post database.Post) ReturnedPost {
	return ReturnedPost{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Url:         post.Url,
		Title:       post.Title,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func DatabasePostssToReturnedPosts(posts []database.Post) []ReturnedPost {
	returnedPosts := make([]ReturnedPost, len(posts))

	for i, post := range posts {
		returnedPosts[i] = DatabasePostToReturnedPost(post)
	}

	return returnedPosts
}

func DatabaseFeedsToReturnedFeeds(feeds []database.Feed) []ReturnedFeed {
	returnedFeeds := make([]ReturnedFeed, len(feeds))

	for i, feed := range feeds {
		returnedFeeds[i] = DatabaseFeedToReturnedFeed(feed)
	}

	return returnedFeeds
}

func DatabaseFeedFollowsToReturnedFollows(feedFollows []database.FeedFollow) []ReturnedFeedFollow {
	returnedFeedFollows := make([]ReturnedFeedFollow, len(feedFollows))

	for i, feedFollow := range feedFollows {
		returnedFeedFollows[i] = DatabaseFeedFollowToResponseFeedFollow(feedFollow)
	}

	return returnedFeedFollows
}

func (apiCfg *apiConfig) GetFeedByURL(r http.Request, url string) (database.Feed, error) {
	feed, err := apiCfg.DB.GetFeedByUrl(r.Context(), url)

	if err != nil {
		return database.Feed{}, errors.New("feed not found")
	}

	return feed, nil
}
