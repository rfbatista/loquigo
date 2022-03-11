package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
)

func NewHoldComponent(c pool.Component) runner.RunnerComponent {
	return HoldComponent{Data: HoldComponentData{FlowId: c.Data.FlowId, StepId: c.Data.StepId}}
}

type HoldComponent struct {
	Data HoldComponentData
}

type HoldComponentData struct {
	FlowId string
	StepId string
}

func (h HoldComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return botMessages, &pool.Stop{StepId: h.Data.StepId, FlowId: h.Data.FlowId}, nil
}
