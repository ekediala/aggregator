package main

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request) {

	respondWithError(w, http.StatusBadRequest, "Something went wrong. Please try again later.")

}
