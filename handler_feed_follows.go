package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ekediala/aggregator/internal/database"
	"github.com/go-chi/chi/v5"
)

func (apiCfg *apiConfig) handleGetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting user feed follows %v", err))
	}

	response := DatabaseFeedFollowsToReturnedFollows(feedFollows)

	respondWithJSON(w, http.StatusOK, ResponseStructure{
		Message: "OK",
		Data:    response,
	})
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")

	if feedFollowIdStr == "" {
		respondWithError(w, 400, "Invalid feed follow id. Please check your url")
	}

	feedFollowId, err := strconv.Atoi(feedFollowIdStr)

	if err != nil {
		respondWithError(w, 400, "Invalid feed follow ID")
	}

	log.Printf("feedfollowId %v", feedFollowId)

	rowsAffected, err := apiCfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{ID: int64(feedFollowId), UserID: user.ID})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error unfollowing feed %v", err))
	}

	respondWithJSON(w, http.StatusAccepted, struct {
		Data string `json:"data"`
	}{
		Data: fmt.Sprintf("%v rows deleted", rowsAffected),
	})
}

func (apiCfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID int64 `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, "JSON malformed")
		return
	}

	if params.FeedID < 1 {
		respondWithError(w, 400, "Invalid ID")
		return
	}

	feedFollow, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not follow feed: %v", err))
		return
	}

	type Response struct {
		Data ReturnedFeedFollow `json:"data"`
	}

	respondWithJSON(w, http.StatusCreated, Response{
		Data: DatabaseFeedFollowToResponseFeedFollow(feedFollow),
	})
}
