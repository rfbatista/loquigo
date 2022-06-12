package dialog

import "loquigo/engine/message"

type Event struct {
	Id      string
	UserID  string
	payload message.Message
}
