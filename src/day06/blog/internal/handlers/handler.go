package handlers

import (
	"day06/blog/internal/config"
	"day06/blog/internal/database"
	"day06/blog/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	articleService *service.ArticlesService
}

func NewHandler(q *database.Queries, cfg *config.Config) *Handler {
	return &Handler{
		articleService: service.NewArticlesService(q, cfg),
	}
}

func (h *Handler) RegisterHTTPEndpoints(r chi.Router) {
	r.Get("/{page}", h.getArticles)
	r.Get("/{id}", h.getArticle)

	r.Get("/admin", h.getLoginPage)
	r.Post("/admin/login", h.loginAdmin)

	r.Post("/admin/post", h.createArticle)
}
