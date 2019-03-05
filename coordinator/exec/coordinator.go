package main

import (
	"github.com/distributed-go/coordinator"
	"time"
)

func main() {
	ea := coordinator.NewEventAggregator()
	_ = coordinator.NewDatabaseConsumer(ea)

	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	for range time.Tick(time.Minute) {

	}
}
