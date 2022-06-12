package dialog

import "loquigo/engine/message"

type TextData struct {
	Body string
}

type ActionData struct {
	ActionMessage message.Message
}

type Action struct {
	id   string
	Type string
	Data ActionData
}
