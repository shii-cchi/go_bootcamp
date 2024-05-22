package app

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

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
		_, err := http.Get(fmt.Sprintf("http://%s:%d/", host, port))

		if err != nil {
			fmt.Printf("Database at %s:%d stopped\n", host, port)
			return
		}
	}
}
