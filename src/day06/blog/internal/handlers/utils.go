package handlers

import (
	"day06/blog/internal/database"
	"encoding/json"
	"log"
	"net/http"
)

type HomePageData struct {
	Articles    []database.Article
	MaxPage     int64
	CurrentPage int64
	PrevPage    int64
	NextPage    int64
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON responce: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}
