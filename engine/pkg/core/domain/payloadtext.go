package domain

type TextPayload struct {
	Type    string `json:"type"`
	Payload string `json:"text"`
}

func CreateTextPayload(s string) TextPayload {
	return TextPayload{string(Text), s}
}
