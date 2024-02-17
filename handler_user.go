package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ekediala/aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 422, fmt.Sprintf("Error parsing JSON %v", r.Body))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), params.Name)

	if err != nil {
		respondWithError(w, 500, "Internal Server Error")
		return
	}

	userToReturn := DatabaseUserToReturnedUser(user)

	respondWithJSON(w, http.StatusCreated, map[string]ReturnedUser{"user": userToReturn})

}

func (apiCfg *apiConfig) handleGetOneUser(w http.ResponseWriter, r *http.Request, user database.User) {

	userToReturn := DatabaseUserToReturnedUser(user)

	respondWithJSON(w, http.StatusOK, map[string]ReturnedUser{"user": userToReturn})
}
