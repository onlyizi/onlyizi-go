package redis

import "github.com/redis/go-redis/v9"

var client *redis.Client

func Client() *redis.Client {
	if client == nil {
		panic("redis not initialized")
	}
	return client
}