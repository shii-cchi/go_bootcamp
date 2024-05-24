package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

type ResponseData struct {
	Code        int      `json:"code"`
	RequestType string   `json:"request_type"`
	Error       string   `json:"error,omitempty"`
	ItemData    ItemData `json:"item_data"`
}

type ItemData struct {
	Name string `json:"name,omitempty"`
}

type Heartbeat struct {
	NodesList         []NodeSummary `json:"nodes_list"`
	ReplicationFactor int           `json:"replication_factor"`
}

type NodeSummary struct {
	Port int    `json:"port"`
	Role string `json:"role"`
}

var heartbeatFromLeader Heartbeat
var port int

func RunClient() {
	host := setupFlags()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		heartbeat(host)
	}()

	makeRequest(host)

	wg.Wait()
}

func setupFlags() string {
	var host string

	flag.StringVar(&host, "H", "", "the host address")
	flag.IntVar(&port, "P", 0, "the port of leader node")

	flag.Parse()

	if host == "" || port == 0 {
		flag.Usage()
		log.Fatal("error: host and/or port are not specified")
	}

	if host != "127.0.0.1" {
		log.Fatal("wrong host address")
	}

	return host
}

func heartbeat(host string) {
	connected := false
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		res, err := http.Get(fmt.Sprintf("http://%s:%d/ping", host, port))

		if err != nil {
			if len(heartbeatFromLeader.NodesList) == 0 {
				log.Fatalf("Node on port %d is not a leader", port)
			}

			fmt.Printf("Leader on port %d is dead\nConnecting to follower on port %d\n", port, heartbeatFromLeader.NodesList[1].Port)

			port = heartbeatFromLeader.NodesList[1].Port
			connected = false
			continue
		}

		err = json.NewDecoder(res.Body).Decode(&heartbeatFromLeader)
		if err != nil {
			log.Fatalf("error decoding heartbeat from leader: %s", err.Error())
		}

		fmt.Println(heartbeatFromLeader.NodesList)

		if !connected {
			fmt.Printf("Connected to a database at %s:%d\n", host, port)
			fmt.Println("Known nodes:")
			for _, node := range heartbeatFromLeader.NodesList {
				fmt.Printf("%s:%d %s\n", host, node.Port, node.Role)
			}
			connected = true
		}
	}
}

func makeRequest(host string) {
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

		res, err := http.Post(fmt.Sprintf("http://%s:%d/", host, port), "application/json", bytes.NewReader(reqJson))

		if err != nil {
			fmt.Println("error sending request:", err)
			break
		}

		handleResponse(res)
	}
}

func handleResponse(res *http.Response) {
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
				fmt.Println("Created" + fmt.Sprintf(" (%d replicas)", heartbeatFromLeader.ReplicationFactor))
			} else {
				fmt.Println("Updated" + fmt.Sprintf(" (%d replicas)", heartbeatFromLeader.ReplicationFactor))
			}

		case "GET":
			fmt.Println(resBody.ItemData)
		case "DELETE":
			fmt.Println("Deleted" + fmt.Sprintf(" (%d replicas)", heartbeatFromLeader.ReplicationFactor))
		}
	}
}
