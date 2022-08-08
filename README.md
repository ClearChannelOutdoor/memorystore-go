# memorystore-go
memorystore-go is a Go module that makes it easy to connect, read, and write to GCP's managed MemoryStore service.

## Usage
Google Cloud's MemoryStore currently supports Redis and Memcache.

## MemoryStore
The `MemoryStore` interface will require Redis and Memcache implementations to look like this:
```go
type MemoryStore interface {
	Delete(key string) error
	Health() error
	Get(key string) (string, error)
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
