package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/templatepool"
)

type HoldComponent struct {
	Component
	Data HoldComponentData
}

type HoldComponentData struct {
	FlowId string
	StepId string
}

func (h HoldComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *templatepool.Stop, *templatepool.GoTo) {
	return botMessages, &templatepool.Stop{StepId: h.StepId, FlowId: h.FlowId}, nil
}
