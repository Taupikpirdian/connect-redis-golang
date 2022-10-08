package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitRedisClient() *redis.Client {
	fmt.Println("Go Redis Connection Test")

	var (
		opts *redis.Options
	)

	opts = &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	}

	client := redis.NewClient(opts)

	return client
}
