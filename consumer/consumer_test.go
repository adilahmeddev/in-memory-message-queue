package consumer

import (
	mq "events-notification"
	"events-notification/helpers"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestConsumer(t *testing.T) {
	t.Run("can consume messages from a message queue", func(t *testing.T) {
		t.Log("given we have a message queue with messages in it")
		messageQueue := mq.NewMemoryQueue()
		message := helpers.RandomMessage()

		assert.NoError(t, messageQueue.Add(message))

		t.Log("we can consume it")
		consumer := NewConsumer(messageQueue)

		gotMessage, err := consumer.Get()
		assert.NoError(t, err)
		assert.Equal(t, message, gotMessage)
	})
}
