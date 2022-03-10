package runner

import (
	"loquigo/engine/src/core/domain"
)

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}
type BotRepository interface {
	FindBotBegin(botId string) (string, error)
}

type FlowRepository interface {
	FindStepBeginIdFromFlow(flowId string) (string, error)
}
type StepRepository interface {
	FindByFlowIdAndStepId(flowId string, stepId string) (RunnerStep, error)
}

type ComponentRepository interface {
	FindByFlowIdAndStepId(flowId string, stepId string) ([]RunnerComponent, error)
}
