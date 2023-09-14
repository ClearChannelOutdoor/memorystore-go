# memorystore-go

memorystore-go is a Go module that makes it easy to connect, read, and write to GCP's managed MemoryStore service.

## Usage

Google Cloud's MemoryStore currently supports Redis and Memcache.

## Unit Tests

You'll need Redis running on your local machine to run unit tests.
To run them, navigate to the `/tests/` directory and run

```bash
REDIS_ADDRESS=localhost REDIS_PORT=55000 REDIS_PASSWORD=put_your_password_here go test
```

## MemoryStore

The `MemoryStore` interface will require Redis and Memcache implementations to look like this:

```go
type MemoryStore interface {
	Delete(key string) error
	Get(key string) (string, error)
	Health() error
	IncrByFloat(key string, value float64) error
	Set(key string, value interface{}, expiration time.Duration) error
}
```

## Redis

### Create a new Redis Store

```go
import memorystore_go "github.com/clearchanneloutdoor/memorystore-go"

func main() {
	config := memorystore_go.RedisConfig {
	    Address: "localhost",
		CACertFile: "file/to/cacert",
		Port: "6379",
		Password: "password",
    }

	redis, err := memorystore_go.NewRedis(config)
	if err != nil {
	    // Handle error
    }
}
```

### Health Check

Health checks can be configured to ensure your service has a working connection to MemoryStore.

```go
func main() {
	if err := redis.Health(); err != nil {
	    // Handle error
    }
}
```

### Get

```go
value, err := redis.Get("key")
```
### IncrByFloat

```go
if err := redis.IncrByFloat("key", 1.0); err != nil {
		// Handle error
}
```

### Set

```go
var value interface{}
if err := redis.Set("key", value, 0); err != nil {
    // Handle error
}
```

### Delete

```go
if err := redis.Delete("key"); err != nil {
    // Handle error
}
```

## Memcache

### Create a new Memcache Store

**_Not Supported Yet_**
