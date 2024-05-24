package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"team01/internal/server/repository"
	"team01/internal/server/service"
	"time"
)

func AllRequestsHandler(store *repository.Store, isLeader bool) http.HandlerFunc {
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

		if (res.Code == http.StatusCreated || (res.Code == http.StatusOK && (res.RequestType == "SET" || res.RequestType == "DELETE"))) && isLeader {
			for _, node := range NodesSummaryList {
				err = makeReplication(node, body)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}
			}
		}

		respondWithJSON(w, res.Code, res)
	}
}

func makeReplication(node NodeSummary, body []byte) error {
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
			fmt.Printf("Successful replication for node on port %d\n", node.Port)
		}

	}

	return nil
}

func HeartbeatFromFollowersHandler(w http.ResponseWriter, r *http.Request) {
	var node Node

	err := json.NewDecoder(r.Body).Decode(&node)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error decoding request body: %s", err.Error()))
		return
	}

	if len(heartbeatFollower.NodesList) != 0 {
		NodesSummaryList = heartbeatFollower.NodesList
		heartbeatFollower = Heartbeat{}
		fmt.Println("NodesList:", NodesSummaryList)
	}

	if isNewNode(node) {
		NodesSummaryList = append(NodesSummaryList, node.NodeSummary)
		NodesList = append(NodesList, node)

		fmt.Printf("the node on port %d has been registered\n", node.NodeSummary.Port)
		fmt.Println("NodesList:", NodesSummaryList)
	} else {
		updateLastActive(node)
	}

	heartbeatLeader := Heartbeat{
		NodesList:         NodesSummaryList,
		ReplicationFactor: REPLICATION_FACTOR,
	}

	respondWithJSON(w, http.StatusOK, heartbeatLeader)
}

func isNewNode(node Node) bool {
	if len(NodesSummaryList) == 0 {
		return true
	}

	for _, n := range NodesSummaryList {
		if n.Port == node.NodeSummary.Port {
			return false
		}
	}

	return true
}

func updateLastActive(node Node) {
	for i, n := range NodesList {
		if n.NodeSummary.Port == node.NodeSummary.Port {
			NodesList[i].LastActive = node.LastActive
			break
		}
	}
}

func CheckFollowers() {
	for i, node := range NodesList {
		if node.NodeSummary.Role == "Leader" {
			continue
		}

		if time.Since(node.LastActive) > HEARTBEAT_TIMEOUT {
			fmt.Printf("Node on port %d is dead\n", node.NodeSummary.Port)
			NodesList = append(NodesList[:i], NodesList[i+1:]...)
			NodesSummaryList = append(NodesSummaryList[:i], NodesSummaryList[i+1:]...)
			fmt.Println("NodesList:", NodesSummaryList)
		}
	}
}

func HeartbeatFromClientHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, Heartbeat{
		NodesList:         NodesSummaryList,
		ReplicationFactor: REPLICATION_FACTOR,
	})
}
