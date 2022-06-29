package consumer

import (
	mq "events-notification"
	"events-notification/messages"
)

type Consumer struct {
	queue mq.Queue
}

func NewConsumer(queue mq.Queue) *Consumer {
	return &Consumer{queue: queue}
}

func (c *Consumer) Get() (messages.Message, error) {
	return c.queue.GetLastMessage()
}
