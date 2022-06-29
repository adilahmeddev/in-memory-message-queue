package helpers

import (
	"events-notification/messages"
	"math/rand"
)

func RandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 15)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomMessage() messages.Message {
	return messages.NewMessage(RandomString())
}
