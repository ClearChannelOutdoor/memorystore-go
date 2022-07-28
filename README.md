# memorystore-go
memorystore-go is a Go module that makes it easy to connect, read, and write to GCP's managed MemoryStore service.

## Usage
Google Cloud's MemoryStore currently supports Redis and Memcache.

## Redis
### Create a new Redis Store
```go
import memorystore_go "github.com/clearchanneloutdoor/pubsub-go"

func main() {
	config := memorystore_go.RedisConfig {
	    Address: "",
		Port: "",
		Password: "",
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

## Memcache
### Create a new Memcache Store
**_Not Supported Yet_**
