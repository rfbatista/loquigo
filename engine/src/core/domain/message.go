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

type ErrorMessage struct {
	ID string `json:"id"`
}

type ImagePayload struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
type QuestionMessage struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
}

type TextPayload struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func CreateTextPayload(s string) TextPayload {
	return TextPayload{string(Text), s}
}
