package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"team01/internal/server/service"
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

func AllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	var reqString RequestString

	err := json.NewDecoder(r.Body).Decode(&reqString)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error decoding req body: %s", err.Error()))
	}

	reqData, err := service.ParseRequest(reqString.DbRequest)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing request string: %s", err.Error()))
	}

	res := service.DoRequest(reqData)

	if res.Error != "" {
		respondWithError(w, res.Code, fmt.Sprintf("error execute db request: %s", res.Error))
	}

	respondWithJSON(w, res.Code, nil)
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
