package consumer

import (
	"context"
	mq "events-notification"
	"events-notification/messages"
	"fmt"
)

type Consumer struct {
	queue    mq.Queue
	messages chan messages.Message
}

func NewConsumer(queue mq.Queue) *Consumer {
	lastMessage := make(chan messages.Message, 1)
	_ = queue.Subscribe(lastMessage)
	return &Consumer{queue: queue, messages: lastMessage}
}

func (c *Consumer) Get(ctx context.Context) (messages.Message, error) {
	select {
	case <-ctx.Done():
		return messages.Message{}, fmt.Errorf("get message error: context deadline exceeded")
	case msg := <-c.messages:

		return msg, nil
	}
}
