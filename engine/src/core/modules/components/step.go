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
	var goTo *pool.GoTo
	var stop *pool.Stop
	var newMessages []domain.Message = messages
	for _, component := range s.Components {
		newMessages, stop, goTo = component.Run(message, context, newMessages)

		if stop != nil {
			return newMessages, stop, nil
		}
		if goTo != nil {
			return newMessages, nil, goTo
		}
	}
	return []domain.Message{}, nil, nil
}

func (s Step) AddComponents(components []runner.RunnerComponent) (runner.RunnerStep, error) {
	return Step{Components: components}, nil
}
