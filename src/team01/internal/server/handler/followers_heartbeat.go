package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"team01/internal/server/config"
	"time"
)

func DoHeartbeat(cfg *config.ServerConfig, cluster *Cluster) {
	newCluster := NewCluster()

	data, err := json.Marshal(Node{Port: cfg.CurrentPort, Role: "Follower", LastActive: time.Now()})
	if err != nil {
		log.Fatal("error marshaling node information")
	}

	res, err := http.Post(fmt.Sprintf("http://127.0.0.1:%d/ping", cfg.LeaderPort), "application/json", bytes.NewReader(data))

	if err != nil {
		if cluster.isEmpty() {
			log.Fatal("The leader was not launched")
		}

		cfg.LeaderPort = cluster.makeNewLeader()
		cluster.PrintNodesList()
		return
	}

	err = json.NewDecoder(res.Body).Decode(&newCluster)
	if err != nil {
		log.Fatal("error decoding response body from leader")
	}

	if newCluster.isEmpty() {
		log.Fatal("cluster is full, current node can't be added")
	}

	if !cluster.IsEqual(newCluster) {
		cluster.update(newCluster)
		cluster.PrintNodesList()
	}

	res.Body.Close()
}
