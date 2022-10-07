package main

import (
	"fmt"
	memorystore_go "github.com/clearchanneloutdoor/memorystore-go"
	// "sync"
)

func main() {
	fmt.Println("Start")

	client, _ := memorystore_go.NewRedisMock()

	client.Health()
	fmt.Println("Done")
}
