package handlers

import (
	"day06/ex01-02/internal/database"
	"day06/ex01-02/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	articleService *service.ArticlesService
}

func NewHandler(q *database.Queries) *Handler {
	return &Handler{
		articleService: service.NewArticlesService(q),
	}
}

func (h *Handler) RegisterHTTPEndpoints(r chi.Router) {
	r.Get("/", h.getArticles)
	r.Get("/", h.getArticle)
	r.Post("/", h.createArticle)
}
