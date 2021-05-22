package redis

import (
	"github.com/go-redis/redis/v8"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(cache *redis.Client) *redisRepo {
	return &redisRepo{
		client: cache,
	}
}
