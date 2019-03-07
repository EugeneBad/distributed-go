package coordinator

import (
	"bytes"
	"encoding/gob"
	"github.com/distributed-go/dto"
	"github.com/distributed-go/monitoring"
	"github.com/distributed-go/qutils"
	"github.com/streadway/amqp"
)

type MetricConsumer struct {
	er      EventRaiser
	conn    *amqp.Connection
	ch      *amqp.Channel
	queue   *amqp.Queue
	sources []string
}

func NewMetricConsumer(er EventRaiser) *DatabaseConsumer {
	dc := DatabaseConsumer{
		er: er,
	}
	dc.conn, dc.ch = qutils.GetChannel(url)
	dc.queue = qutils.GetQueue(qutils.LiveReadingsQueue, dc.ch, false)

	dc.er.AddListener("DataSourceDiscovered", func(eventData interface{}) {
		dc.SubscribeToDataEvent(eventData.(string))
	})
	return &dc
}

func (mc *MetricConsumer) SubscribeToDataEvent(eventName string) {
	for _, v := range mc.sources {
		if v == eventName {
			return
		}
	}

	callback := mc.callbackGenerator()

	mc.er.AddListener("MessageReceived_"+eventName, callback)
}

func (mc *MetricConsumer) callbackGenerator() func(interface{}) {
	buf := new(bytes.Buffer)
	rc := monitoring.NewReadingCounter()

	return func(eventData interface{}) {
		ed := eventData.(EventData)

		sm := dto.SensorMessage{
			Name:      ed.Name,
			Value:     ed.Value,
			Timestamp: ed.Timestamp,
		}
		buf.Reset()
		enc := gob.NewEncoder(buf)
		_ = enc.Encode(sm)

		msg := amqp.Publishing{
			Body: buf.Bytes(),
		}

		err := mc.ch.Publish("", qutils.LiveReadingsQueue, false, false, msg)

		if err == nil {
			rc.Increment("coordinator", ed.Name)
		}

	}
}
