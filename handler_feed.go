package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ekediala/aggregator/internal/database"
)

type ResponseStructure struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func handlerGetOneFeed(w http.ResponseWriter, r *http.Request, user database.User){

// }

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())

	if err != nil {
		respondWithError(w, 500, "Internal Server Error")
		return
	}

	type Response struct {
		Feeds []ReturnedFeed `json:"feeds"`
	}

	respondWithJSON(w, http.StatusOK, Response{Feeds: DatabaseFeedsToReturnedFeeds(feeds)})
}

func (apiCfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, "JSON malformed")
		return
	}

	if strings.TrimSpace(params.Name) == "" {
		respondWithError(w, 400, "Feed name is required")
		return
	}

	if strings.TrimSpace(params.Url) == "" {
		respondWithError(w, 400, "Feed URL is required")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:   params.Name,
		Url:    params.Url,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not save feed: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Printf("Feed follow could not be created: %v", err)
	}

	respondWithJSON(w, http.StatusCreated, ResponseStructure{
		Message: "Feed created successfully",
		Data: struct {
			Feed       ReturnedFeed       `json:"feed"`
			FeedFollow ReturnedFeedFollow `json:"feed_follow"`
		}{
			Feed:       DatabaseFeedToReturnedFeed(feed),
			FeedFollow: DatabaseFeedFollowToResponseFeedFollow(feedFollow),
		},
	})

}
