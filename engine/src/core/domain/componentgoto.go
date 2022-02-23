package domain

type GotoComponent struct {
	FlowId string
	StepId string
}

func NewGoTo(flowId string, stepId string) GotoComponent {
	return GotoComponent{FlowId: flowId, StepId: stepId}
}

func (g GotoComponent) Run(m Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo) {
	return botMessages, nil, &GoTo{FlowId: g.FlowId, StepId: g.StepId}
}
