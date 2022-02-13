package domain

type TextComponent struct {
	FlowId string
	StepId string
	text   string
}

func (t TextComponent) Run(m Message, u UserContext) ([]Message, *Stop, error) {
	var message Message = CreateTextMessage(t.text)
	return []Message{message}, nil, nil
}
