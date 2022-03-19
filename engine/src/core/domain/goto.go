package domain

func NewGotoComponent(c Component) GotoComponent {
	return GotoComponent{Data: ComponentData{GroupId: c.Data.GroupId, NodeId: c.Data.NodeId}}
}

type GotoComponent struct {
	Data ComponentData `json:"data"`
}

type GoToComponentData struct {
	GroupId string
	NodeId  string
}

func (g GotoComponent) Run(m Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo) {
	return botMessages, nil, &GoTo{GroupId: g.Data.GroupId, NodeId: g.Data.NodeId}
}
