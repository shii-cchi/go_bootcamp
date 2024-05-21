package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"team01/internal/server/handler"
)

func RunServer() {
	r := chi.NewRouter()

	r.Post("/", handler.AllRequestsHandler)

	log.Fatal(http.ListenAndServe(":8888", r))
}
