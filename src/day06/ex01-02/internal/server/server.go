package server

import (
	"database/sql"
	"day06/ex01-02/internal/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	httpServer  *http.Server
	httpHandler *http.Handler
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

	services := service.NewServices(queries, cfg)

	handler := handlers.NewHandler(services)

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
