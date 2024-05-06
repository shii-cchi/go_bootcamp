package handlers

import (
	"day06/blog/internal/handlers/dto"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
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

	articles, err := h.articleService.GetArticles(r.Context(), int64(page))

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

	article, err := h.articleService.GetArticle(r.Context(), int32(id))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error getting article with id %d: %v", id, err))
		return
	}

	respondWithJSON(w, http.StatusOK, article)
}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {
	newArticle := new(dto.ArticleDto)
	err := json.NewDecoder(r.Body).Decode(&newArticle)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	article, err := h.articleService.CreateArticle(r.Context(), newArticle.Title, newArticle.Content)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error creating article: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, article)
}

func (h *Handler) getLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("login page").Parse("")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing html template: %v", err))
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error executing template: %v", err))
		return
	}
}

func (h *Handler) loginAdmin(w http.ResponseWriter, r *http.Request) {
	adminCredentials := new(dto.AdminDto)
	err := json.NewDecoder(r.Body).Decode(&adminCredentials)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	err = h.articleService.LoginAdmin(adminCredentials.Login, adminCredentials.Password)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}
