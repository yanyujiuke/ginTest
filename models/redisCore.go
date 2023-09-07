package models

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	ctx1    = context.Background()
	RedisDb *redis.Client
)

func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(ctx1).Result()
	if err != nil {
		println(err)
	}
}
