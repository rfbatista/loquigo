package domain

type TextComponent struct {
	text string
}

func NewText(text string) TextComponent {
	return TextComponent{text: text}
}

func (t TextComponent) Run(m Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo) {
	var message Message = NewTextMessage(t.text)
	messages := append(botMessages, message)
	return messages, nil, nil
}
