package aredis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func Get[T any](ctx context.Context, key string, bind *T) (bool, error) {
	result, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, err
	}
	if err != nil {
		return true, err
	}

	if err := json.Unmarshal([]byte(result), bind); err != nil {
		return true, err
	}

	return true, err
}
