package templatepool

type Component struct {
	ID       string        `json:"id"`
	FlowId   string        `json:"flow_id" bson:"flow_id"`
	StepId   string        `json:"step_id" bson:"step_id"`
	Type     string        `json:"type" bson:"type"`
	Data     ComponentData `json:"data" bson:"data"`
	Sequence int           `json:"sequence"`
}

type ComponentData struct {
	Text   string `json:"text" bson:"text,omitempty"`
	FlowId string `json:"flow_id" bson:"flow_id,omitempty"`
	StepId string `json:"step_id" bson:"step_id,omitempty"`
}

func NewComponent(c Component) IComponent {
	var dc IComponent
	switch c.Type {
	case "text":
		dc = NewText(c.Data.Text)
	case "hold":
		dc = NewHold(c.Data.FlowId, c.Data.StepId)
	case "goto":
		dc = NewGoTo(c.Data.FlowId, c.Data.StepId)
	}
	return dc
}
