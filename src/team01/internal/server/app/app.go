package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"team01/internal/server/db"
	"team01/internal/server/handler"
)

func RunServer() {
	database := db.NewDatabase()

	r := chi.NewRouter()

	r.Post("/", handler.AllRequestsHandler(database))
	r.Get("/", handler.HeartbeatHandler)

	log.Fatal(http.ListenAndServe(":8888", r))
}
