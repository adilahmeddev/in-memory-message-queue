package events_notification

import (
	"events-notification/helpers"
	"events-notification/messages"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestQueue(t *testing.T) {
	var (
		messageQueue = NewMemoryQueue()
		message      = helpers.RandomMessage()
	)

	t.Run("I am able to add a message to the queue", func(t *testing.T) {
		err := messageQueue.Add(message)
		assert.NoError(t, err)
	})
	t.Run("I am able to read the latest message from the queue", func(t *testing.T) {
		latestMessage := helpers.RandomMessage()
		err := messageQueue.Add(latestMessage)
		assert.NoError(t, err)

		gotMessage, err := messageQueue.GetLastMessage()
		assert.NoError(t, err)

		assert.Equal(t, latestMessage, gotMessage)
		t.Run("and the message is removed from the queue", func(t *testing.T) {
			gotMessage2, err := messageQueue.GetLastMessage()
			assert.NoError(t, err)

			assert.NotEqual(t, latestMessage, gotMessage2)
			assert.Equal(t, message, gotMessage2)
		})
	})
	t.Run("I am able to get the length of the message queue", func(t *testing.T) {
		mq := NewMemoryQueue()

		count, err := mq.GetMessageCount()
		assert.NoError(t, err)
		assert.Equal(t, 0, count)

		assert.NoError(t, mq.Add(helpers.RandomMessage()))

		count, err = mq.GetMessageCount()
		assert.NoError(t, err)
		assert.Equal(t, 1, count)

		_, err = mq.GetLastMessage()
		assert.NoError(t, err)

		count, err = mq.GetMessageCount()
		assert.NoError(t, err)
		assert.Equal(t, 0, count)

	})

	t.Run("can subscribe to the message queue", func(t *testing.T) {
		mq := NewMemoryQueue()
		listener := make(chan messages.Message, 1)
		mq.Subscribe(listener)
	})

}
