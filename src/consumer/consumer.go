package consumer

import (
	"context"
	"events-notification/models/messages"
	mq "events-notification/src"

	"fmt"
)

type Consumer struct {
	queue    mq.Queue
	messages chan messages.Message
	version  int
}

func NewConsumer(queue mq.Queue) *Consumer {
	lastMessage := make(chan messages.Message, 1)
	_ = queue.Subscribe(lastMessage)
	return &Consumer{queue: queue, messages: lastMessage}
}

func (c *Consumer) Get(ctx context.Context) (messageList []messages.Message, err error) {
	select {
	case <-ctx.Done():
		if c.version < c.queue.GetVersion() {
			messageList, err = c.syncOldMessages(messageList)
			if err != nil {
				return nil, err
			}
			return messageList, nil
		}
		return []messages.Message{}, fmt.Errorf("get message error: context deadline exceeded")
	case msg := <-c.messages:
		messageList = []messages.Message{msg}
		c.version++
		if c.version < c.queue.GetVersion() {
			messageList, err = c.syncOldMessages(messageList)
			if err != nil {
				return nil, err
			}
		}
		return messageList, nil
	}
}

func (c *Consumer) syncOldMessages(messageList []messages.Message) ([]messages.Message, error) {
	for c.version < c.queue.GetVersion() {
		message, err := c.queue.GetMessage(c.version)
		if err != nil {
			return []messages.Message{}, err
		}
		messageList = append([]messages.Message{message}, messageList...)
		c.version++
	}
	return messageList, nil
}
