package main

import (
	"fmt"
	"github.com/distributed-go/coordinator"
)

var dc *coordinator.DatabaseConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	//dc = coordinator.NewDatabaseConsumer(ea)

	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	_, _ = fmt.Scanln(&a)
}
