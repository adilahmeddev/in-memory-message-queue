package events_notification

import (
	"events-notification/helpers"
	"events-notification/messages"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestQueue(t *testing.T) {
	messageQueue := NewMemoryQueue()
	message := messages.NewMessage("Hello world")
	t.Run("I am able to add a message to the queue", func(t *testing.T) {
		err := messageQueue.Add(message)
		assert.NoError(t, err)
	})
	t.Run("I am able to read the latest message from the queue", func(t *testing.T) {
		err := messageQueue.Add(helpers.RandomMessage())
		assert.NoError(t, err)

		gotMessage, err := messageQueue.GetLastMessage()
		assert.NoError(t, err)

		assert.Equal(t, message, gotMessage)
	})

}

func TestProducer(t *testing.T) {
	//t.Run("when I create a message with my producer, I expect it to be sent to its destination", func(t *testing.T) {
	//	messageQueue := NewMessageQueue()
	//	producer := NewProducer(messageQueue)
	//})
}
