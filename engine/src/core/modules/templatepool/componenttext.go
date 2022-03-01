package templatepool

import "loquigo/engine/src/core/domain"

type TextComponent struct {
	text string
}

func NewText(text string) TextComponent {
	return TextComponent{text: text}
}

func (t TextComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *Stop, *GoTo) {
	var message domain.Message = domain.NewTextMessage(t.text)
	messages := append(botMessages, message)
	return messages, nil, nil
}
