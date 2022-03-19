package nodes

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/components"
)

func NewNodeRunner(node domain.Node) RunnerNode {
	return NodeRunner{}
}

type NodeRunner struct {
	ID         string
	BotId      string `json:"bot_id"`
	GroupId    string `json:"group_id"`
	Name       string `json:"name"`
	Components []components.RunnerComponent
}

func (s NodeRunner) Run(message domain.Message, context domain.UserContext, messages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	var goTo *domain.GoTo
	var stop *domain.Stop
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

func (s NodeRunner) AddComponents(components []components.RunnerComponent) (RunnerNode, error) {
	return NodeRunner{Components: components}, nil
}

func (s NodeRunner) IsValid() bool {
	if len(s.Components) > 0 {
		return true
	} else {
		return false
	}
}
