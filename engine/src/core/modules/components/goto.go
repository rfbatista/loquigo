package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
)

func NewGotoComponent(c pool.Component) runner.RunnerComponent {
	return GotoComponent{Data: GoToComponentData{FlowId: c.Data.FlowId, StepId: c.Data.StepId}}
}

type GotoComponent struct {
	Data GoToComponentData
}

type GoToComponentData struct {
	FlowId string
	StepId string
}

func (g GotoComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return botMessages, nil, &pool.GoTo{FlowId: g.Data.FlowId, StepId: g.Data.StepId}
}
