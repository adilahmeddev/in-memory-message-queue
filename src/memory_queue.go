package src

import (
	"events-notification/models/messages"
	"fmt"
	"github.com/alecthomas/repr"
)

type Queue interface {
	Add(message messages.Message) error
	GetLastMessage() (messages.Message, error)
	GetMessageCount() (int, error)
	Subscribe(listener chan messages.Message) error
	GetVersion() int
	GetMessage(i int) (messages.Message, error)
}

type MemoryQueue struct {
	listeners []chan messages.Message
	messages  []messages.Message
	history   []messages.Message
	version   int
}

func (m *MemoryQueue) GetVersion() int {
	return m.version
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
func (m *MemoryQueue) GetMessage(i int) (messages.Message, error) {
	if i > len(m.messages)-1 {
		return messages.Message{}, fmt.Errorf("message number %v is out of bounds", i)
	}
	return m.history[i], nil
}
func (m *MemoryQueue) Add(message messages.Message) error {
	m.history = append(m.history, message)
	m.messages = append(m.messages, message)
	m.version++
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
