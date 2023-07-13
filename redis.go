package memorystore_go

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"time"
)

const (
	localAddress = "localhost"
)

type Redis struct {
	client RedisClient
	ctx    context.Context
}

type RedisConfig struct {
	Address    string
	CACertFile string
	Port       string
	Password   string
}

func NewRedis(config RedisConfig) (Redis, error) {
	var tlsConfig *tls.Config
	if config.Address != localAddress && config.CACertFile != "" {
		caCert, err := ioutil.ReadFile(config.CACertFile)
		if err != nil {
			return Redis{}, err
		}

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

	r := Redis{
		client: client,
		ctx:    ctx,
	}

	return r, r.Health()
}

func NewRedisMock(client *redis.Client) Redis {
	ctx := context.Background()

	r := Redis{
		client: client,
		ctx:    ctx,
	}

	return r
}

func (r Redis) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r Redis) Health() error {
	return r.client.Ping(r.ctx).Err()
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
