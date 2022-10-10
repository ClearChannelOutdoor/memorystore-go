package memorystore_tests

import (
	"encoding/json"
	"errors"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisMockGet(t *testing.T) {
	db, mock := redismock.NewClientMock()

	client := memorystore_go.NewRedisMock(db)

	mock.ExpectGet("get-key").RedisNil()

	value, _ := client.Get("get-key")

	assert.Equal(t, value, "")
	mock.ClearExpect()
}

func TestRedisMockWrite(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := memorystore_go.NewRedisMock(db)
	valueBytes, encodingError := json.Marshal("set-value")
	if encodingError != nil {
		t.Error(encodingError)
	}
	mock.ExpectSet("set-key", valueBytes, 10*time.Minute).SetErr(errors.New("FAIL"))

	err := client.Set("set-key", "set-value", 10*time.Minute)

	if err == nil || err.Error() != "FAIL" {
		t.Error("Wrong Error Recieved")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	mock.ClearExpect()
}
