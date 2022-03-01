package templatepool

import "loquigo/engine/src/core/domain"

type HoldComponent struct {
	FlowId string
	StepId string
}

func NewHold(flowId string, stepId string) HoldComponent {
	return HoldComponent{FlowId: flowId, StepId: stepId}
}

func (h HoldComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *Stop, *GoTo) {
	return botMessages, &Stop{StepId: h.StepId, FlowId: h.FlowId}, nil
}
