package service

import (
	"context"
	"day06/blog/internal/config"
	"day06/blog/internal/database"
	"errors"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"math"
	"time"
)

const limitArticles = 3

type ArticlesService struct {
	queries *database.Queries
	cfg     *config.Config
}

func NewArticlesService(q *database.Queries, cfg *config.Config) *ArticlesService {
	return &ArticlesService{
		queries: q,
		cfg:     cfg,
	}
}

func (s ArticlesService) GetArticles(ctx context.Context, page int) ([]database.GetArticlesRow, int, error) {
	articlesCount, err := s.queries.GetArticlesCount(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("error getting articles count: %v", err)
	}

	pageCount := int(math.Ceil(float64(articlesCount) / float64(limitArticles)))

	if page != 1 && page > pageCount {
		return nil, 0, fmt.Errorf("out of range page value: %v", err)
	}

	articles, err := s.queries.GetArticles(ctx, database.GetArticlesParams{
		Limit:  limitArticles,
		Offset: int32((page - 1) * limitArticles),
	})

	if err != nil {
		return nil, 0, fmt.Errorf("error getting articles: %v", err)
	}

	return articles, pageCount, nil
}

func (s ArticlesService) GetArticle(ctx context.Context, id int32) (database.GetArticleRow, error) {
	article, err := s.queries.GetArticle(ctx, id)

	if err != nil {
		return database.GetArticleRow{}, err
	}

	return database.GetArticleRow{
		ID:      article.ID,
		Title:   article.Title,
		Content: string(blackfriday.Run([]byte(article.Content))),
	}, nil
}

func (s ArticlesService) CreateArticle(ctx context.Context, title, content string) (database.Article, error) {
	article, err := s.queries.CreateArticle(ctx, database.CreateArticleParams{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		return database.Article{}, err
	}

	return article, nil
}

func (s ArticlesService) LoginAdmin(login, password string) error {
	if login != s.cfg.AdminLogin || password != s.cfg.AdminPassword {
		return errors.New("wrong admin credentials")
	}

	return nil
}
