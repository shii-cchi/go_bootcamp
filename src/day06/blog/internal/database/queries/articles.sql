-- name: GetArticles :many
SELECT id, title, content, created_at FROM articles
LIMIT $1 OFFSET $2;

-- name: GetArticle :one
SELECT title, content, created_at FROM articles
WHERE id = $1;

-- name: CreateArticle :one
INSERT INTO articles (title, content, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetArticlesCount :one
SELECT COUNT(*) AS article_count FROM articles;