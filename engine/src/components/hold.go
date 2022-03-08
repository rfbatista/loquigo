package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type HoldComponent struct {
	Component
	Data HoldComponentData
}

type HoldComponentData struct {
	FlowId string
	StepId string
}

func (h HoldComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return botMessages, &pool.Stop{StepId: h.StepId, FlowId: h.FlowId}, nil
}
