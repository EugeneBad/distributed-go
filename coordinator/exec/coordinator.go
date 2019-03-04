package main

import (
	"fmt"
	"github.com/distributed-go/coordinator"
)

func main() {
	ea := coordinator.NewEventAggregator()
	_ = coordinator.NewDatabaseConsumer(ea)

	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	_, _ = fmt.Scanln(&a)
}
