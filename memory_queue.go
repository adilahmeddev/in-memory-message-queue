package events_notification

import "events-notification/messages"

type Queue interface {
	Add(message messages.Message) error
}

type MemoryQueue struct {
}

func (m MemoryQueue) Add(message messages.Message) error {
	//TODO implement me
	panic("implement me")
}

func NewMemoryQueue() Queue {
	return &MemoryQueue{}
}

