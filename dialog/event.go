package dialog

import "loquigo/engine/message"

type Event struct {
	Id      string
	BotID   string
	UserID  string
	payload message.Message
}
