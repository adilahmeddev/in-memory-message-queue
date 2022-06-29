package consumer

import (
	"context"
	mq "events-notification"
	"events-notification/helpers"
	"events-notification/messages"
	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	t.Run("given multiple consumers", func(t *testing.T) {
		t.Run("they can both consume the same message from the message queue", func(t *testing.T) {
			var (
				ctx, cancelFunc = context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
				messageQueue    = mq.NewMemoryQueue()
				consumerB       = NewConsumer(messageQueue)
				consumerA       = NewConsumer(messageQueue)
				message         = helpers.RandomMessage()
			)

			assert.NoError(t, messageQueue.Add(message))

			gotA, err := consumerA.Get(ctx)
			assert.NoError(t, err)
			assert.Equal(t, gotA, message)

			gotB, err := consumerB.Get(ctx)
			assert.NoError(t, err)
			assert.Equal(t, gotB, message)

			cancelFunc()

		})

	})
}

func contains(t *testing.T, messages []messages.Message, message messages.Message) {
	found := false
	for _, m := range messages {
		if m == message {
			found = true
		}
	}
	if !found {
		t.Errorf("message list %s does not contain message %s", repr.String(messages), message)
	}
}
