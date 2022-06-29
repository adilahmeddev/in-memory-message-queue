package producer

import (
	mq "events-notification"
	"events-notification/messages"
)

type Producer struct {
	queue mq.Queue
}

func (p Producer) Add(message messages.Message) error {
	return p.queue.Add(message)
}

func NewProducer(queue mq.Queue) *Producer {
	return &Producer{
		queue: queue,
	}
}
