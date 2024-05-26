package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"team01/internal/client/config"
)

func MakeRequest(cfg *config.ClientConfig, heartbeat *Heartbeat) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		reqStr := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("error input:", err)
			break
		}

		req := RequestString{DbRequest: reqStr}

		reqJson, err := json.Marshal(req)

		if err != nil {
			fmt.Println("error marshalling request:", err)
			break
		}

		select {
		case newPort := <-cfg.PortChan:
			cfg.Port = newPort
		default:
		}

		res, err := http.Post(fmt.Sprintf("http://%s:%d/", cfg.Host, cfg.Port), "application/json", bytes.NewReader(reqJson))

		if err != nil {
			fmt.Println("error sending request:", err)
			break
		}

		handleResponse(res, heartbeat)
	}
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
		switch resBody.RequestType {
		case "SET":
			if resBody.Code == http.StatusCreated {
				fmt.Println("Created" + fmt.Sprintf(" (%d replicas)", heartbeat.ReplicationFactor))
			} else {
				fmt.Println("Updated" + fmt.Sprintf(" (%d replicas)", heartbeat.ReplicationFactor))
			}

		case "GET":
			fmt.Println(resBody.ItemData)
		case "DELETE":
			fmt.Println("Deleted" + fmt.Sprintf(" (%d replicas)", heartbeat.ReplicationFactor))
		}
	}
}
