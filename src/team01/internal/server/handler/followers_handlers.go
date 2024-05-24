package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

var heartbeatFollower Heartbeat

func DoHeartbeat(port, leaderPort int) {
	node := Node{NodeSummary: NodeSummary{Port: port, Role: "Follower"}}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	var newHeartbeatFollower Heartbeat

	for range ticker.C {
		node.LastActive = time.Now()

		data, err := json.Marshal(node)
		if err != nil {
			log.Fatal("error marshaling node information")
		}

		res, err := http.Post(fmt.Sprintf("http://127.0.0.1:%d/ping", leaderPort), "application/json", bytes.NewReader(data))

		if err != nil {
			if len(heartbeatFollower.NodesList) == 0 {
				log.Fatal("The leader was not launched")
			}

			leaderPort = makeNewLeader()

			if leaderPort == port {
				break
			}

			continue
		}

		err = json.NewDecoder(res.Body).Decode(&newHeartbeatFollower)
		if err != nil {
			log.Fatal("error decoding response body from leader")
		}

		if !reflect.DeepEqual(newHeartbeatFollower, heartbeatFollower) {
			heartbeatFollower = newHeartbeatFollower
			fmt.Println(heartbeatFollower.NodesList)
		}

		res.Body.Close()
	}
}

func makeNewLeader() int {
	fmt.Println("Leader is dead")

	heartbeatFollower.NodesList = heartbeatFollower.NodesList[1:]

	heartbeatFollower.NodesList[0].Role = "Leader"
	leaderPort := heartbeatFollower.NodesList[0].Port

	fmt.Printf("New leader is node on port %d\n", leaderPort)

	return leaderPort
}
