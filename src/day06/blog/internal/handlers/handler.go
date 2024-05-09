package handlers

import (
	"day06/blog/internal/config"
	"day06/blog/internal/database"
	"day06/blog/internal/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang.org/x/time/rate"
	"net/http"
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
	r.Use(rateLimit(100))

	fs := http.FileServer(http.Dir("./blog_frontend"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", h.getArticles)
	r.Get("/{id}", h.getArticle)

	r.Get("/admin", h.getLoginPage)
	r.Post("/admin/login", h.loginAdmin)

	r.Get("/admin/post", h.getNewPostPage)
	r.Post("/admin/post", h.createArticle)
}

func rateLimit(requests int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		limiter := rate.NewLimiter(rate.Limit(requests), requests)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if limiter.Allow() == false {
				fmt.Fprint(w, "<p style='color:red;'>429 Too Many Requests</p>")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
