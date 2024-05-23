package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"team01/internal/server/handler"
	"team01/internal/server/repository"
)

func RunServer() {
	database := repository.NewStore()

	r := chi.NewRouter()

	r.Post("/", handler.AllRequestsHandler(database))
	r.Get("/ping", handler.HeartbeatHandler)

	log.Fatal(http.ListenAndServe(":8888", r))
}
