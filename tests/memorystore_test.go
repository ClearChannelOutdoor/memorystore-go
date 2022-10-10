package memorystore_tests

import (
	// "fmt"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	// "sync"
	"encoding/json"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestData struct {
	Value string
}

func TestRedisMockGet(t *testing.T) {

	//var d = TestData{
	//	Value: "woosh",
	//}

	db, mock := redismock.NewClientMock()

	client := memorystore_go.NewRedisMock(db)

	mock.ExpectGet("get-key").RedisNil()

	value, _ := client.Get("get-key")

	assert.Equal(t, value, "")
	mock.ClearExpect()

	// r.client.Ping(r.ctx).Err()
}

func TestRedisMockWrite(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := memorystore_go.NewRedisMock(db)
	valueBytes, encodingError := json.Marshal("set-value")
	if encodingError != nil {
		t.Error(encodingError)
	}
	mock.ExpectSet("set-key", valueBytes, 10*time.Minute)

	client.Set("set-key", "set-value", 10*time.Minute)

	//if err == nil || err.Error() != "FAIL" {
	//	t.Error("Wrong Error Recieved")
	//}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	mock.ClearExpect()
}
