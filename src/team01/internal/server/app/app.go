package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"team01/internal/server/config"
	"team01/internal/server/handler"
	"team01/internal/server/repository"
)

func RunServer() {
	database := repository.NewStore()

	cfg := config.SetupFlags()

	cluster := handler.NewCluster()

	if cfg.CurrentPort == cfg.LeaderPort {
		cluster.AppendNode(handler.Node{Port: cfg.CurrentPort, Role: "Leader"})
	}

	go cluster.Monitor(&cfg)

	r := chi.NewRouter()

	r.Post("/", handler.AllRequestsHandler(database, &cfg, cluster))
	r.Post("/ping", handler.HeartbeatFromFollowersHandler(cluster))
	r.Get("/ping", handler.HeartbeatFromClientHandler(cluster))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.CurrentPort), r))
}
