package redis_repository

import (
	"context"
	"fmt"
	"try/connect-redis-golang/domain/entity"
	"try/connect-redis-golang/redis_repository/mapper"

	"github.com/go-redis/redis/v8"
)

type RepoArticleRedis struct {
	Conn *redis.Client
}

func NewRepoArticleRedisInteractor(Conn *redis.Client) *RepoArticleRedis {
	return &RepoArticleRedis{Conn: Conn}
}

func (repo *RepoArticleRedis) GetAttributeArticleByKode(ctx context.Context, kodeArticle string) (*entity.Article, error) {
	var (
		checkErr error
	)

	// data, checkErr = repo.Conn.HGet(ctx, kodeArticle, mapper.MapGetRedisField()).Result()
	data, checkErr := repo.Conn.HGetAll(ctx, mapper.MapGetRedisField()).Result()

	if len(data) == 0 || checkErr == redis.Nil {
		fmt.Println("Redis is Empty")
		// return nil, checkErr // kenapa ketika return error, dia malah error
		return nil, nil
	}

	fmt.Println("... Yeah found data in Redis")
	// dataArticle, err := mapper.MapFromJsonStringToDomainArticle(data)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (repo *RepoArticleRedis) StoreOrUpdateData(ctx context.Context, data *entity.Article) error {
	/*
		CLI Redis:
		HGET "XXXX-25524000" "data_article"
	*/
	fmt.Println("... Store Data To Redis")
	_, err := repo.Conn.HSet(ctx, mapper.MapGetKeyValueRedis(data), mapper.MapGetRedisField(), mapper.MapSetBukuToString(data)).Result()
	if err != nil {
		return err
	}

	return nil
}
