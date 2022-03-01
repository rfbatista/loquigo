package templatepool

import "loquigo/engine/src/core/domain"

type GotoComponent struct {
	FlowId string
	StepId string
}

func NewGoTo(flowId string, stepId string) GotoComponent {
	return GotoComponent{FlowId: flowId, StepId: stepId}
}

func (g GotoComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *Stop, *GoTo) {
	return botMessages, nil, &GoTo{FlowId: g.FlowId, StepId: g.StepId}
}
