package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"team01/internal/client/config"
	"time"
)

func DoHeartbeat(cfg *config.ClientConfig, heartbeat *Heartbeat) {
	connected := false
	ticker := time.NewTicker(heartbeatTick)

	for range ticker.C {
		res, err := http.Get(fmt.Sprintf("http://%s:%d/ping", cfg.Host, cfg.Port))

		if err != nil {
			handleHeartbeatError(cfg, heartbeat)
			connected = false
			continue
		}

		err = json.NewDecoder(res.Body).Decode(&heartbeat)
		if err != nil {
			log.Fatalf("error decoding heartbeat from leader: %s", err.Error())
		}

		if !connected {
			printConnectionMessage(cfg, heartbeat)
			connected = true
		}
	}
}

func handleHeartbeatError(cfg *config.ClientConfig, heartbeat *Heartbeat) {
	if len(heartbeat.NodesList) == 0 {
		log.Fatalf("Node on port %d is not a leader", cfg.Port)
	}

	if len(heartbeat.NodesList) == 1 {
		log.Fatal("There are no running nodes")
	}

	fmt.Printf("Leader on port %d is dead\nConnecting to follower on port %d\n", cfg.Port, heartbeat.NodesList[1].Port)

	cfg.Port = heartbeat.NodesList[1].Port
}

func printConnectionMessage(cfg *config.ClientConfig, heartbeat *Heartbeat) {
	fmt.Printf("Connected to a database at %s:%d\n", cfg.Host, cfg.Port)
	fmt.Println("Known nodes:")
	for _, node := range heartbeat.NodesList {
		fmt.Printf("%s:%d %s\n", cfg.Host, node.Port, node.Role)
	}

	if len(heartbeat.NodesList) < heartbeat.ReplicationFactor {
		fmt.Printf("WARNING: cluster size (%d) is smaller than a replication factor (%d)!\n", len(heartbeat.NodesList), heartbeat.ReplicationFactor)
	}
}
