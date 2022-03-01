package templatepool

import "loquigo/engine/src/core/domain"

type IStep = func(message domain.Message, context domain.UserContext, messages []domain.Message) []IComponent

func NewStep(ID string, FlowID string, Name string) Step {
	return Step{
		ID:     ID,
		FlowId: FlowID,
		Name:   Name,
	}
}

type Step struct {
	ID         string `json:"id"`
	FlowId     string `json:"flow_id"`
	Name       string `json:"name"`
	Components []IComponent
}

func (s Step) Run(message domain.Message, context domain.UserContext, messages []domain.Message) []IComponent {
	return s.Components
}
