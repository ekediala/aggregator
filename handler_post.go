package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ekediala/aggregator/internal/database"
)

func (apiCfg apiConfig) handlerGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	// we can't parse limit from query param, default to 20
	if limit == 0 || err != nil {
		limit = 20
	}

	feedFollows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not find feed for this user: %v", err))
		return
	}

	feedIDs := make([]int32, len(feedFollows))

	for i, feed := range feedFollows {
		feedIDs[i] = int32(feed.ID)
	}

	posts, err := apiCfg.DB.GetPostsByFeedFollow(r.Context(), database.GetPostsByFeedFollowParams{
		Column1: feedIDs,
		Limit:   int32(limit),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not find feed for this user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseStructure{
		Data:    DatabasePostssToReturnedPosts(posts),
		Message: "Posts found",
	})

}
