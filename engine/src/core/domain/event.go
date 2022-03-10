package domain

type Event struct {
	Bot       Bot     `json:"bot"`
	User      User    `json:"user"`
	Message   Message `json:"message"`
	Timestamp string  `json:"timestamp"`
}

type Bot struct {
	ID      string `json:"id" bson:"id"`
	BeginId string `bson:"begin_id"`
}
