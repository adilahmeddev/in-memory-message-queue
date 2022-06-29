package events_notification

import (
	"events-notification/messages"
	"github.com/alecthomas/repr"
)

type Queue interface {
	Add(message messages.Message) error
	GetLastMessage() (messages.Message, error)
	GetMessageCount() (int, error)
	Subscribe(listener chan messages.Message) error
}

type MemoryQueue struct {
	listeners []chan messages.Message
	messages  []messages.Message
	history   []messages.Message
}

func (m *MemoryQueue) Subscribe(listener chan messages.Message) error {
	m.listeners = append(m.listeners, listener)
	return nil
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
	m.history = append(m.history, message)
	m.messages = append(m.messages, message)
	go func(messageParam messages.Message) {
		repr.Println("listeners: ", m.listeners)
		repr.Println("adding ", message)
		for _, listener := range m.listeners {
			listener <- message
		}
	}(message)
	return nil
}

func NewMemoryQueue() Queue {
	return &MemoryQueue{}
}
