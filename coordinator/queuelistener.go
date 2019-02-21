package coordinator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/distributed-go/dto"
	"github.com/distributed-go/qutils"
	"github.com/streadway/amqp"
)

const url = "amqp://guest@localhost:5672"

type QueueListener struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
}

func NewQueueListener() *QueueListener {
	ql := QueueListener{}

	ql.conn, ql.ch = qutils.GetChannel(url)
	ql.sources = map[string]<-chan amqp.Delivery{}
	return &ql
}

func (ql *QueueListener) ListenForNewSource() {
	q := qutils.GetQueue("", ql.ch)

	_ = ql.ch.QueueBind(
		q.Name,
		"",
		"amq.fanout",
		false,
		nil)

	msgs, _ := ql.ch.Consume(q.Name, "", true, false, false, false, nil)

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body),
			"",
			true,
			false,
			false,
			false,
			nil)
		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)

		sd := new(dto.SensorMessage)

		_ = d.Decode(sd)

		fmt.Printf("Received message: %v\n", sd)
	}
}
