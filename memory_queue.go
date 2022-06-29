package events_notification

import (
	"events-notification/messages"
)

type Queue interface {
	Add(message messages.Message) error
	GetLastMessage() (messages.Message, error)
	GetMessageCount() (int, error)
}

type MemoryQueue struct {
	messages []messages.Message
}

func (m *MemoryQueue) GetMessageCount() (int, error) {
	return len(m.messages), nil
}

func (m *MemoryQueue) GetLastMessage() (messages.Message, error) {
	message := m.messages[len(m.messages)-1]
	m.messages = m.messages[:len(m.messages)-1]
	return message, nil
}

func (m *MemoryQueue) Add(message messages.Message) error {
	m.messages = append(m.messages, message)
	return nil
}

func NewMemoryQueue() Queue {
	return &MemoryQueue{}
}
