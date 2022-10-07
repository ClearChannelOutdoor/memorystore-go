package memorystore_go

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type MemoryStore interface {
	Delete(key string) error
	Health() error
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type RedisClient interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
