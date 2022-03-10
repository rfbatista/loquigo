package components

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
)

type Step struct {
	ID         string
	Components []runner.RunnerComponent
}

func (s Step) Run(message domain.Message, context domain.UserContext, messages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo) {
	return []domain.Message{}, nil, nil
}

func (s Step) AddComponents(components []runner.RunnerComponent) error {
	s.Components = components
	return nil
}
