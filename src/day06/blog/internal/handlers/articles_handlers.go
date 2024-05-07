package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) getArticles(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page <= 0 {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка: страница не найдена</p>")
		return
	}

	articles, maxPage, err := h.articleService.GetArticles(r.Context(), int64(page))

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка: страница не найдена</p>")
		return
	}

	homePageData := HomePageData{
		Articles:    articles,
		MaxPage:     maxPage,
		CurrentPage: int64(page),
		PrevPage:    int64(page - 1),
		NextPage:    int64(page + 1),
	}

	tmpl := template.Must(template.ParseFiles("blog_frontend/html/main_page.html"))

	if err := tmpl.Execute(w, homePageData); err != nil {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка</p>")
		return
	}
}

func (h *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка: статья не найдена</p>")
		return
	}

	article, err := h.articleService.GetArticle(r.Context(), int32(id))

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка: статья не найдена</p>")
		return
	}

	tmpl := template.Must(template.ParseFiles("blog_frontend/html/article_page.html"))

	if err := tmpl.Execute(w, article); err != nil {
		fmt.Fprint(w, "<p style='color:red;'>Ошибка</p>")
		return
	}
}

func (h *Handler) getLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("blog_frontend/html/login_page.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error executing template: %v", err))
		return
	}

	errorMessage := r.URL.Query().Get("error")

	if errorMessage != "" {
		fmt.Fprintf(w, "<p style='color:red;'>Ошибка: %s</p>", errorMessage)
	}
}

func (h *Handler) loginAdmin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing login data: %v", err))
		return
	}

	adminLogin := r.Form.Get("username")
	adminPassword := r.Form.Get("password")

	err = h.articleService.LoginAdmin(adminLogin, adminPassword)

	if err != nil {
		http.Redirect(w, r, "/admin?error=Неверное имя пользователя или пароль", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin/post", http.StatusSeeOther)
}

func (h *Handler) getNewPostPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("blog_frontend/html/admin_panel.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error executing template: %v", err))
		return
	}
}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing login data: %v", err))
		return
	}

	articleTitle := r.Form.Get("title")
	articleContent := r.Form.Get("content")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	article, err := h.articleService.CreateArticle(r.Context(), articleTitle, articleContent)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error creating article: %v", err))
		return
	}

	articleId := article.ID

	http.Redirect(w, r, fmt.Sprintf("/%d", articleId), http.StatusSeeOther)
}
