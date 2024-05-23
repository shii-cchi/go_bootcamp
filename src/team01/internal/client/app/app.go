package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
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

func RunClient() {
	host, port := setupFlags()

	if host == "" || port == 0 {
		flag.Usage()
		return
	}

	if host != "127.0.0.1" {
		fmt.Println("wrong host address")
		return
	}

	fmt.Printf("Connected to a database of Warehouse 13 at %s:%d\n", host, port)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		heartbeat(host, port)
	}()

	makeRequest(host, port)

	wg.Wait()
}

func setupFlags() (string, int) {
	var host string
	var port int

	flag.StringVar(&host, "H", "", "the host address")
	flag.IntVar(&port, "P", 0, "the port")

	flag.Parse()

	return host, port
}

func heartbeat(host string, port int) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		_, err := http.Get(fmt.Sprintf("http://%s:%d/ping", host, port))

		if err != nil {
			fmt.Printf("Database at %s:%d stopped\n", host, port)
			return
		}
	}
}

func makeRequest(host string, port int) {
	var req RequestString
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		reqStr := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("error input:", err)
		}

		req.DbRequest = reqStr

		reqJson, _ := json.Marshal(req)

		res, err := http.Post(fmt.Sprintf("http://%s:%d/", host, port), "application/json", bytes.NewReader(reqJson))

		if err != nil {
			fmt.Println("error sending request:", err)
		}

		var resBody ResponseData

		err = json.NewDecoder(res.Body).Decode(&resBody)

		if err != nil {
			fmt.Println("error decoding response body:", err)
		}

		if resBody.Error != "" {
			fmt.Println(resBody.Error)
		} else {
			switch resBody.RequestType {
			case "SET":
				if resBody.Code == http.StatusCreated {
					fmt.Println("Created")
				} else {
					fmt.Println("Updated")
				}

			case "GET":
				fmt.Println(resBody.ItemData)
			case "DELETE":
				fmt.Println("Deleted")
			}
		}

	}
}
