package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
)

func NewTextComponent(c pool.Component) runner.RunnerComponent {
	return TextComponent{Data: TextComponentData{Text: c.Data.Text}}
}

type TextComponent struct {
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
