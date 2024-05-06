package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) getArticles(w http.ResponseWriter, r *http.Request) {
	pageStr := chi.URLParam(r, "page")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page <= 0 {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid page value '%s': %v", pageStr, err))
		return
	}

	articles, err := h.articleService.GetArticles(context.Background(), int64(page))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error getting articles on page %d: %v", page, err))
		return
	}

	respondWithJSON(w, http.StatusOK, articles)
}

func (h *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "invalid id value")
		return
	}

	article, err := h.articleService.GetArticle(context.Background(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error getting article with id %d: %v", id, err))
		return
	}

	respondWithJSON(w, http.StatusOK, article)
}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {

}
