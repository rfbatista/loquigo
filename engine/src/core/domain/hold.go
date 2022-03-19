package domain

func NewHoldComponent(c Component) HoldComponent {
	return HoldComponent{Data: ComponentData{GroupId: c.Data.GroupId, NodeId: c.Data.NodeId}}
}

type HoldComponent struct {
	Data ComponentData
}

type HoldComponentData struct {
	GroupId string
	NodeId  string
}

func (h HoldComponent) Run(m Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo) {
	return botMessages, &Stop{NodeId: h.Data.NodeId, GroupId: h.Data.GroupId}, nil
}
