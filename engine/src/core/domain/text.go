package domain

func NewTextComponent(c Component) TextComponent {
	return TextComponent{Data: TextComponentData{Text: c.Data.Text}}
}

type TextComponent struct {
	Data TextComponentData
}

type TextComponentData struct {
	Text string
}

func (t TextComponent) Run(m Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo) {
	var message Message = NewTextMessage(t.Data.Text)

	messages := append(botMessages, message)
	return messages, nil, nil
}
