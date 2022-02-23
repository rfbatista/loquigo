package templatepool

import (
	"loquigo/engine/src/core/domain"
)

type FlowHash = map[string]StepHash

type Flow interface {
	Name() string
	Steps(string) Step
}

type StepHash = map[string]Step

type Step = func(message domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component
