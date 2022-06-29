package events_notification

import (
	"events-notification/messages"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestQueue(t *testing.T) {
	messageQueue := NewMemoryQueue()
	t.Run("I am able to add messages to the queue", func(t *testing.T) {
		message := messages.NewMessage("Hello world")
		err := messageQueue.Add(message)
		assert.NoError(t, err)
	})
}

func TestProducer(t *testing.T) {
	//t.Run("when I create a message with my producer, I expect it to be sent to its destination", func(t *testing.T) {
	//	messageQueue := NewMessageQueue()
	//	producer := NewProducer(messageQueue)
	//})
}
