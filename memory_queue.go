package events_notification

import (
	"events-notification/messages"
)

type Queue interface {
	Add(message messages.Message) error
	GetLastMessage() (messages.Message, error)
}

type MemoryQueue struct {
	messages []messages.Message
}

func (m *MemoryQueue) GetLastMessage() (messages.Message, error) {
	return m.messages[len(m.messages)-1], nil
}

func (m *MemoryQueue) Add(message messages.Message) error {
	m.messages = append(m.messages, message)
	return nil
}

func NewMemoryQueue() Queue {
	return &MemoryQueue{}
}
