package main

import (
	"net/http"

	"github.com/ekediala/aggregator/internal/auth"
	"github.com/ekediala/aggregator/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByAPiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, "User not found")
			return
		}

		handler(w, r, user)
	}
}
