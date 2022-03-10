package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type GotoComponent struct {
	Component
	Data HoldComponentData
}

type GoToComponentData struct {
	FlowId string
	StepId string
}

func (g GotoComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return botMessages, nil, &pool.GoTo{FlowId: g.Data.FlowId, StepId: g.Data.StepId}
}
