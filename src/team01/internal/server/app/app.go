package app

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"team01/internal/server/handler"
	"team01/internal/server/repository"
	"time"
)

func RunServer() {
	database := repository.NewStore()

	port, leaderPort := setupFlags()

	if port == leaderPort {
		node := handler.Node{
			NodeSummary: handler.NodeSummary{
				Port: port,
				Role: "Leader",
			},
		}

		handler.NodesSummaryList = append(handler.NodesSummaryList, node.NodeSummary)
		handler.NodesList = append(handler.NodesList, node)
		fmt.Println("NodesList:", handler.NodesSummaryList)

		go func() {
			for {
				handler.CheckFollowers()
				time.Sleep(5 * time.Second)
			}
		}()

	} else {
		go handler.DoHeartbeat(port, leaderPort)
	}

	r := chi.NewRouter()

	r.Post("/", handler.AllRequestsHandler(database, port == leaderPort))
	r.Post("/ping", handler.HeartbeatFromFollowersHandler)
	r.Get("/ping", handler.HeartbeatFromClientHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func setupFlags() (int, int) {
	var port int
	var leaderPort int

	flag.IntVar(&port, "P", 0, "the port of this node")
	flag.IntVar(&leaderPort, "L", 0, "the port of leader node")

	flag.Parse()

	if port == 0 || leaderPort == 0 {
		flag.Usage()
		log.Fatal("error: ports are not specified")
	}

	return port, leaderPort
}
