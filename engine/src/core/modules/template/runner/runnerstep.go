package runner

import "loquigo/engine/src/core/domain"

type RunnerStep struct {
	ID         string `json:"id"`
	FlowId     string `json:"flow_id"`
	Name       string `json:"name"`
	Components []IComponent
}

func (s RunnerStep) Run(message domain.Message, context domain.UserContext, messages []domain.Message) []IComponent {
	return s.Components
}
