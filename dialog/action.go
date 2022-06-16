package dialog

import (
	"loquigo/engine/message"
)

var collection = "action"

type ActionType int16

const (
	UndefinedType ActionType = iota
	TextMessage
	ImageMessage
)

func (m ActionType) String() string {
	switch m {
	case TextMessage:
		return "text_message"
	case ImageMessage:
		return "image_message"
	}
	return "unknown"
}

type TextData struct {
	Body string
}

type ActionData struct {
	ActionMessage message.Message
}

type Action struct {
	id   string
	Type ActionType
	Data ActionData
}

func (a Action) Run(event Event) *message.Message {
	return &a.Data.ActionMessage
}
