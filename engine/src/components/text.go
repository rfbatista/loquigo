package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type TextComponent struct {
	Component
	Data TextComponentData
}

type TextComponentData struct {
	Text string
}

func (t TextComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	var message domain.Message = domain.NewTextMessage(t.Data.Text)
	messages := append(botMessages, message)
	return messages, nil, nil
}
