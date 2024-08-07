package app

import (
	"day06/blog/internal/server"
	"github.com/go-chi/chi/v5"
	"log"
)

func Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := chi.NewRouter()

	srv, err := server.NewServer(r)

	if err != nil {
		log.Fatal(err)
	}

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
