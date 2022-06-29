package messages

type Message struct {
	Contents string
}

func NewMessage(contents string) Message {
	return Message{Contents: contents}
}
