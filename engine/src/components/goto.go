package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/templatepool"
)

type GotoComponent struct {
	Component
	Data HoldComponentData
}

type GoToComponentData struct {
	FlowId string
	StepId string
}

func (g GotoComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *templatepool.Stop, *templatepool.GoTo) {
	return botMessages, nil, &templatepool.GoTo{FlowId: g.Data.FlowId, StepId: g.Data.StepId}
}
