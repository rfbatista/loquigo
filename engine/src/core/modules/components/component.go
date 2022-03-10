package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
)

type Component struct {
	ID       string        `json:"id"`
	FlowId   string        `json:"flowId"`
	StepId   string        `json:"stepId"`
	Type     string        `json:"type" `
	Data     ComponentData `json:"data"`
	Sequence int           `json:"sequence"`
}
type ComponentData struct{}

func (c Component) Run(userMess domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return botMessages, nil, nil
}

func BuildRunnerComponent(c pool.Component) runner.RunnerComponent {
	return Component{}
}
