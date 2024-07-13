package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"team01/internal/server/config"
	"team01/internal/server/repository"
	"team01/internal/server/service"
)

func AllRequestsHandler(store *repository.Store, cfg *config.ServerConfig, cluster *Cluster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqString RequestString

		body, err := io.ReadAll(r.Body)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error reading request body: %s", err.Error()))
			return
		}

		err = json.Unmarshal(body, &reqString)
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

		if (res.Code == http.StatusCreated || (res.Code == http.StatusOK && (res.RequestType == "SET" || res.RequestType == "DELETE"))) && cfg.CurrentPort == cfg.LeaderPort {
			for _, node := range cluster.NodesList {
				err = MakeReplication(node, body)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}
			}
		}

		respondWithJSON(w, res.Code, res)
	}
}

func MakeReplication(node Node, body []byte) error {
	if node.Role != "Leader" {
		replRes, err := http.Post(fmt.Sprintf("http://127.0.0.1:%d/", node.Port), "application/json", bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("error sending request for replication: %s", err.Error())
		}

		defer replRes.Body.Close()

		var replResBody service.ResponseData

		err = json.NewDecoder(replRes.Body).Decode(&replResBody)
		if err != nil {
			return fmt.Errorf("error decoding res body by replication: %s", err.Error())
		}

		if replResBody.Error == "" {
			log.Printf("Successful replication for node on port %d\n", node.Port)
		} else {
			log.Println(replResBody.Error)
		}

	}

	return nil
}

func HeartbeatFromFollowersHandler(store *repository.Store, cluster *Cluster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		node := Node{}

		err := json.NewDecoder(r.Body).Decode(&node)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error decoding request body: %s", err.Error()))
			return
		}

		if !cluster.isExistNode(node) {
			if cluster.isFull() {
				log.Printf("the maximum number of nodes (%d), a new node with a port=(%d) can't be added\n", ReplicationFactor, node.Port)

				respondWithError(w, http.StatusConflict, "cluster is full")
				return
			} else {
				cluster.AppendNode(node)
				err = cluster.SyncNewNode(node, service.GetStore(store))

				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}
			}
		} else {
			cluster.updateLastActive(node)
		}

		respondWithJSON(w, http.StatusOK, cluster)
	}
}

func HeartbeatFromClientHandler(cluster *Cluster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, Heartbeat{NodesList: cluster.NodesList, ReplicationFactor: ReplicationFactor})
	}
}
