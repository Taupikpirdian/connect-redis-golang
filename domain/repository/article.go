package repository

import (
	"context"
	"try/connect-redis-golang/domain/entity"
)

type ArticleRepository interface {
	Store(ctx context.Context, dataArticle *entity.Article) error
	GetAllData(ctx context.Context) ([]*entity.Article, error)
}

type ArticleRedisRepository interface {
	GetAttributeArticleByKode(ctx context.Context, codeArticle string) (*entity.Article, error)
	StoreOrUpdateData(ctx context.Context, dataArticle *entity.Article) error
}
