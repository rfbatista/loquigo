package domain

type TextPayload struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func CreateTextPayload(s string) TextPayload {
	return TextPayload{string(Text), s}
}
