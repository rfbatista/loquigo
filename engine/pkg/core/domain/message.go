package domain

type MessageType string

const (
	Text  MessageType = "text"
	Image MessageType = "image"
)

type Payload interface{}

type Message struct {
	ID   string      `json:"id"`
	Type MessageType `json:"type"`
	Payload
}

func CreateTextMessage(s string) Message {
	return Message{"", Text, CreateTextPayload(s)}
}
