package domain

type MessageType string

const (
	Text  MessageType = "text"
	Image MessageType = "image"
)

type Payload interface{}

type Message struct {
	ID      string `json:"id"`
	Payload `json:"payload"`
}

func CreateTextMessage(s string) Message {
	return Message{"", CreateTextPayload(s)}
}
