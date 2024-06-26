// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: articles.sql

package database

import (
	"context"
	"time"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (title, content, created_at)
VALUES ($1, $2, $3)
RETURNING id, title, content, created_at
`

type CreateArticleParams struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, createArticle, arg.Title, arg.Content, arg.CreatedAt)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getArticle = `-- name: GetArticle :one
SELECT id, title, content FROM articles
WHERE id = $1
`

type GetArticleRow struct {
	ID      int32
	Title   string
	Content string
}

func (q *Queries) GetArticle(ctx context.Context, id int32) (GetArticleRow, error) {
	row := q.db.QueryRowContext(ctx, getArticle, id)
	var i GetArticleRow
	err := row.Scan(&i.ID, &i.Title, &i.Content)
	return i, err
}

const getArticles = `-- name: GetArticles :many
SELECT id, title, created_at FROM articles
LIMIT $1 OFFSET $2
`

type GetArticlesParams struct {
	Limit  int32
	Offset int32
}

type GetArticlesRow struct {
	ID        int32
	Title     string
	CreatedAt time.Time
}

func (q *Queries) GetArticles(ctx context.Context, arg GetArticlesParams) ([]GetArticlesRow, error) {
	rows, err := q.db.QueryContext(ctx, getArticles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetArticlesRow
	for rows.Next() {
		var i GetArticlesRow
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArticlesCount = `-- name: GetArticlesCount :one
SELECT COUNT(*) AS article_count FROM articles
`

func (q *Queries) GetArticlesCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getArticlesCount)
	var article_count int64
	err := row.Scan(&article_count)
	return article_count, err
}
