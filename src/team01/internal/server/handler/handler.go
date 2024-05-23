package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"team01/internal/server/repository"
	"team01/internal/server/service"
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

func AllRequestsHandler(store *repository.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqString RequestString

		err := json.NewDecoder(r.Body).Decode(&reqString)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error decoding req body: %s", err.Error()))
			return
		}

		reqData, err := service.ParseRequest(reqString.DbRequest)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		res := service.DoRequest(reqData, store)

		respondWithJSON(w, res.Code, res)
	}
}

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "pong")
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
