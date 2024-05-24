package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

type Heartbeat struct {
	NodesList         []NodeSummary `json:"nodes_list"`
	ReplicationFactor int           `json:"replication_factor"`
}

type Node struct {
	NodeSummary NodeSummary `json:"node_summary"`
	LastActive  time.Time   `json:"last_active"`
}

type NodeSummary struct {
	Port int    `json:"port"`
	Role string `json:"role"`
}

const REPLICATION_FACTOR = 2
const HEARTBEAT_TIMEOUT = time.Second * 11

var NodesList []Node
var NodesSummaryList []NodeSummary

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
