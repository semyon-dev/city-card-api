package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var onceRedis sync.Once

func ConnectToRedis() {
	onceRedis.Do(connectToRedis)
}

var Redis *redis.Client

func connectToRedis() {
	url := os.Getenv("REDIS_URI")
	if url == "" {
		log.Println("No set REDIS_URI")
		return
	}
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	Redis = redis.NewClient(opt)
	pong, err := Redis.Ping(context.TODO()).Result()
	fmt.Println(pong, err)
}
