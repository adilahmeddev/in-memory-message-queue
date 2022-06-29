package producer

import (
	"events-notification/models/helpers"
	mq "events-notification/src"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestProducer(t *testing.T) {
	t.Run("when I create a message with my producer, I expect it to be sent to its destination", func(t *testing.T) {
		messageQueue := mq.NewMemoryQueue()
		producer := NewProducer(messageQueue)

		err := producer.Add(helpers.RandomMessage())
		assert.NoError(t, err)
	})
}
