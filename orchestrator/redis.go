package main

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type Redis struct {
	Client *redis.Client
}

func (r *Redis) SetString(key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *Redis) GetString(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
