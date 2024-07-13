package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"team01/internal/client/config"
)

func MakeRequest(cfg *config.ClientConfig, heartbeat *Heartbeat, fr *FailedRequests) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		reqStr := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("error input:", err)
			continue
		}

		req := RequestString{DbRequest: reqStr}

		reqJson, err := json.Marshal(req)

		if err != nil {
			fmt.Println("error marshalling request:", err)
			continue
		}

		res, err := http.Post(fmt.Sprintf("http://%s:%d/", cfg.Host, cfg.Port), "application/json", bytes.NewReader(reqJson))

		if err != nil {
			if len(heartbeat.NodesList) == 0 {
				fmt.Println("Failed to write/read an entry: all nodes are dead")
			}

			fr.failedRequests = append(fr.failedRequests, FailedRequest{RequestString: reqJson})
			continue
		}

		handleResponse(res, heartbeat)
	}
}

func retryFailedRequests(cfg *config.ClientConfig, heartbeat *Heartbeat, fr *FailedRequests) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	for _, failedReq := range fr.failedRequests {
		res, err := http.Post(fmt.Sprintf("http://%s:%d/", cfg.Host, cfg.Port), "application/json", bytes.NewReader(failedReq.RequestString))

		req := strings.TrimRight(strings.ReplaceAll(strings.SplitN(string(failedReq.RequestString), ":", 2)[1], `\`, ``), `}`)

		if err != nil {
			fmt.Printf("Failed to resend for request: %s\n", req)
			continue
		}

		fmt.Printf("Resend for request : %s\n", req)
		handleResponse(res, heartbeat)
	}

	fr.failedRequests = nil
}

func handleResponse(res *http.Response, heartbeat *Heartbeat) {
	defer res.Body.Close()

	var resBody ResponseData

	err := json.NewDecoder(res.Body).Decode(&resBody)

	if err != nil {
		fmt.Println("error decoding response body:", err)
		return
	}

	if resBody.Error != "" {
		fmt.Println(resBody.Error)
	} else {
		countReplicas := heartbeat.ReplicationFactor

		if len(heartbeat.NodesList)-1 < heartbeat.ReplicationFactor {
			countReplicas = len(heartbeat.NodesList) - 1
		}

		switch resBody.RequestType {
		case "SET":
			if resBody.Code == http.StatusCreated {
				fmt.Println("Created" + fmt.Sprintf(" (%d replicas)", countReplicas))
			} else {
				fmt.Println("Updated" + fmt.Sprintf(" (%d replicas)", countReplicas))
			}

		case "GET":
			fmt.Println(resBody.ItemData)
		case "DELETE":
			fmt.Println("Deleted" + fmt.Sprintf(" (%d replicas)", countReplicas))
		}
	}
}
