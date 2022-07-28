package memorystore_go

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	localAddress = "localhost"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

type RedisConfig struct {
	Address  string
	Port     string
	Password string
}

func NewRedis(config RedisConfig) (Redis, error) {
	var tlsConfig *tls.Config
	if config.Address != localAddress {
		// TODO: Read in file
		var caCert []byte

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    caCertPool,
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:      fmt.Sprintf("%s:%s", config.Address, config.Port),
		Password:  config.Password,
		TLSConfig: tlsConfig,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return Redis{}, err
	}

	return Redis{
		client: client,
		ctx:    ctx,
	}, nil
}

func (r Redis) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r Redis) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r Redis) Set(key string, value interface{}, expiration time.Duration) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	result := r.client.Set(r.ctx, key, valueBytes, expiration)

	return result.Err()
}
