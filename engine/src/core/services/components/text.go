package components

import "loquigo/engine/src/core/domain"

func NewTextComponent(c domain.Component) TextComponent {
	return TextComponent{domain.Component{Data: domain.ComponentData{Text: c.Data.Text}}}
}

type TextComponent struct {
	domain.Component
}

func (t TextComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	var message domain.Message = domain.NewTextMessage(t.Data.Text)
	messages := append(botMessages, message)
	return messages, nil, nil
}
