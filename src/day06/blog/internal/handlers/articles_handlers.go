package handlers

import (
	"day06/blog/internal/database"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

type HomePageData struct {
	Articles    []database.GetArticlesRow
	MaxPage     int64
	CurrentPage int64
	PrevPage    int64
	NextPage    int64
}

type ArticleData struct {
	ID      int32
	Title   string
	Content template.HTML
}

func (h *Handler) getArticles(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page <= 0 {
		fmt.Fprint(w, "<p style='color:red;'>404 Page not found</p>")
		return
	}

	articles, maxPage, err := h.articleService.GetArticles(r.Context(), int64(page))

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>404 Page not found</p>")
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
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}
}

func (h *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		fmt.Fprint(w, "<p style='color:red;'>404 Article not found</p>")
		return
	}

	article, err := h.articleService.GetArticle(r.Context(), int32(id))

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>404 Article not found</p>")
		return
	}

	articleData := ArticleData{
		ID:      article.ID,
		Title:   article.Title,
		Content: template.HTML(article.Content),
	}

	tmpl := template.Must(template.ParseFiles("blog_frontend/html/article_page.html"))

	if err := tmpl.Execute(w, articleData); err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}
}

func (h *Handler) getLoginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
		http.Redirect(w, r, "/admin/post", http.StatusSeeOther)
		return
	}

	errorMessage := ""

	if session.Values["authenticated"] != nil && !session.Values["authenticated"].(bool) {
		errorMessage = "Неверный логин или пароль"
	}

	tmpl := template.Must(template.ParseFiles("blog_frontend/html/login_page.html"))

	if err := tmpl.Execute(w, errorMessage); err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}
}

func (h *Handler) loginAdmin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}

	adminLogin := r.Form.Get("username")
	adminPassword := r.Form.Get("password")

	err = h.articleService.LoginAdmin(adminLogin, adminPassword)

	if err != nil {
		session, _ := store.Get(r, "session")
		session.Values["authenticated"] = false
		session.Options.MaxAge = 60
		session.Save(r, w)

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Options.MaxAge = 7200
	session.Save(r, w)

	http.Redirect(w, r, "/admin/post", http.StatusSeeOther)
}

func (h *Handler) getNewPostPage(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["authenticated"] == nil || !session.Values["authenticated"].(bool) {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("blog_frontend/html/admin_panel.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}
}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}

	articleTitle := r.Form.Get("title")
	articleContent := r.Form.Get("content")

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}

	article, err := h.articleService.CreateArticle(r.Context(), articleTitle, articleContent)

	if err != nil {
		fmt.Fprint(w, "<p style='color:red;'>500 InternalServerError</p>")
		return
	}

	articleId := article.ID

	http.Redirect(w, r, fmt.Sprintf("/%d", articleId), http.StatusSeeOther)
}
