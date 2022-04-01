package domain

type Event struct {
	Bot       EventBot `json:"bot"`
	User      User     `json:"user"`
	Message   Message  `json:"message"`
	Timestamp string   `json:"timestamp"`
}

type EventBot struct {
	ID string `json:"id"`
}
