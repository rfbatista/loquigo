package message

type MessageType int16

const (
	UndenfinedMessage MessageType = iota
	TextMessage
	ImageMessage
)

func (m MessageType) String() string {
	switch m {
	case TextMessage:
		return "text"
	case ImageMessage:
		return "image"
	}
	return "unknown"
}

type Message struct {
	Id   string      `json:"id"`
	Type MessageType `json:"type"`
	Text TextData    `json:"text,omitempty"`
}
