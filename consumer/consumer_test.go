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

	t.Run("given multiple consumers", func(t *testing.T) {
		t.Run("they can both consume the same message from the message queue", func(t *testing.T) {
			var (
				messageQueue = mq.NewMemoryQueue()
				consumerB    = NewConsumer(messageQueue)
				consumerA    = NewConsumer(messageQueue)
				message      = helpers.RandomMessage()
			)

			assert.NoError(t, messageQueue.Add(message))

			gotA, err := consumerA.Get()
			assert.NoError(t, err)
			assert.Equal(t, message, gotA)

			gotB, err := consumerB.Get()
			assert.NoError(t, err)
			assert.Equal(t, message, gotB)
		})

	})
}
