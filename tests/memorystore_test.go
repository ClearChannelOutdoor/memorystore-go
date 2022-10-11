package memorystore_tests

import (
	"encoding/json"
	"errors"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

const (
	testKey = "db_test_key"
)

func database() memorystore_go.RedisConfig {
	return memorystore_go.RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}

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
		t.Error("Wrong Error Received")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	mock.ClearExpect()
}

func TestRedisMockDel(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := memorystore_go.NewRedisMock(db)

	mock.ExpectDel("del-key").SetErr(errors.New("FAIL"))

	err := client.Delete("del-key")

	if err == nil || err.Error() != "FAIL" {
		t.Error("Wrong Error Received")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	mock.ClearExpect()
}

func TestRedisSet(t *testing.T) {
	client, err := memorystore_go.NewRedis(database())

	if err != nil {
		t.Error(err)
	}

	setErr := client.Set(testKey, "values", 10*time.Minute)
	assert.Nil(t, setErr)
}

func TestRedisGet(t *testing.T) {
	client, err := memorystore_go.NewRedis(database())

	if err != nil {
		t.Error(err)
	}

	value, _ := client.Get(testKey)

	assert.Equal(t, value, "\"values\"")
}

func TestRedisDel(t *testing.T) {
	client, err := memorystore_go.NewRedis(database())

	if err != nil {
		t.Error(err)
	}

	delErr := client.Delete(testKey)

	assert.Nil(t, delErr)

	value, _ := client.Get(testKey)

	assert.Equal(t, value, "")
}
