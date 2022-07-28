package memorystore_go

import "time"

type MemoryStore interface {
	Delete(key string) error
	Health() error
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}
