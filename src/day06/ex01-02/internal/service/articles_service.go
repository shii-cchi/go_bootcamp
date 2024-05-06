package service

import (
	"context"
	"day06/ex01-02/internal/database"
	"fmt"
)

const limitArticles = 3

type ArticlesService struct {
	queries *database.Queries
}

func NewArticlesService(q *database.Queries) *ArticlesService {
	return &ArticlesService{
		queries: q,
	}
}

func (s ArticlesService) GetArticles(ctx context.Context, page int64) ([]database.GetArticlesRow, error) {
	articlesCount, err := s.queries.GetArticlesCount(ctx)

	if err != nil {
		return nil, fmt.Errorf("error getting articles count: %v", err)
	}

	if page > articlesCount {
		return nil, fmt.Errorf("out of range page value: %v", err)
	}

	articles, err := s.queries.GetArticles(ctx, database.GetArticlesParams{
		Limit:  limitArticles,
		Offset: int32((page - 1) * limitArticles),
	})

	if err != nil {
		return nil, fmt.Errorf("error getting articles: %v", err)
	}

	return articles, nil
}

func (s ArticlesService) GetArticle(ctx context.Context, id int64) (database.GetArticleRow, error) {
	article, err := s.queries.GetArticle(ctx, id)

	if err != nil {
		return database.GetArticleRow{}, err
	}

	return article, nil
}
