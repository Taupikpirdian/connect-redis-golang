package repository

import (
	"context"
	"try/connect-redis-golang/domain/entity"
)

type ArticleRepository interface {
	Store(ctx context.Context, dataArticle *entity.Article) error
	GetAllData(ctx context.Context) ([]*entity.Article, error)
}
