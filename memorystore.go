package memorystore_go

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type MemoryStore interface {
	Delete(key string) error
	Get(key string) (string, error)
	Health() error
	IncrByFloat(key string, value float64) error
	Set(key string, value interface{}, expiration time.Duration) error
}

type RedisClient interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
