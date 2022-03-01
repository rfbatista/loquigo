package domain

type MessageType string

const (
	Text  MessageType = "text"
	Image MessageType = "image"
)

type Payload interface{}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewTextMessage(s string) Message {
	m := CreateTextPayload(s)
	return Message{"", m.Type, m.Text}
}
