package memorystore_tests

import (
	"fmt"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	// "sync"
	"github.com/stretchr/testify/assert"
	"testing"
	// "time"
	"github.com/go-redis/redismock/v8"
)

type TestData struct {
	Value string
}

func TestRedisMockWrite(t *testing.T) {

	//var d = TestData{
	//	Value: "woosh",
	//}

	db, mock := redismock.NewClientMock()

	client := memorystore_go.NewRedisMock(db)

	mock.ExpectGet("get-key").RedisNil()

	value, err := client.Get("get-key")

	fmt.Println(err)

	assert.Equal(t, value, "")

	// r.client.Ping(r.ctx).Err()
}
