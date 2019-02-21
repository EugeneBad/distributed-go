package main

import (
	"fmt"
	"github.com/distributed-go/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	_, _ = fmt.Scanln(&a)
}
