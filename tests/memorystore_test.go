package memorystore_tests

import (
	// "fmt"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	// "sync"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisMockCreate(t *testing.T) {

	client, _ := memorystore_go.NewRedisMock()

	assert.Nil(t, client.Health())
}
