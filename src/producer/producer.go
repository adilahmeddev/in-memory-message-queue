package producer

import (
	"events-notification/models/messages"
	. "events-notification/src"
)

type Producer struct {
	queue Queue
}

func (p Producer) Add(message messages.Message) error {
	return p.queue.Add(message)
}

func NewProducer(queue Queue) *Producer {
	return &Producer{
		queue: queue,
	}
}
