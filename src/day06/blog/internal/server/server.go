package server

import (
	"database/sql"
	"day06/blog/internal/config"
	"day06/blog/internal/database"
	"day06/blog/internal/handlers"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Server struct {
	httpServer  *http.Server
	httpHandler *handlers.Handler
	queries     *database.Queries
}

func NewServer(r chi.Router) (*Server, error) {
	cfg, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("postgres", cfg.DbURI)

	if err != nil {
		return nil, err
	}

	queries := database.New(conn)

	handler := handlers.NewHandler(queries, cfg)

	handler.RegisterHTTPEndpoints(r)

	log.Printf("Server starting on port %s", cfg.Port)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
		},
		httpHandler: handler,
	}, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
