package pool

import (
	"loquigo/engine/src/core/domain"
)

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}

type BotRepository interface {
	FindBeginByBotId(botId string) (string, error)
}
type FlowRepository interface {
	FindByBotId(id string) ([]Flow, error)
	Create(flow Flow) (Flow, error)
	Update(flow Flow) (Flow, error)
	Delete(flow Flow) (Flow, error)
	DeleteByBotID(botId string) error
}
type StepRepository interface {
	FindByFlowId(id string) ([]Step, error)
	FindByIdAndFlowId(flowId string, stepId string) (Step, error)
	FindById(id string) (Step, error)
	Create(step Step) (Step, error)
	Update(step Step) (Step, error)
	Delete(step Step) (Step, error)
	DeleteByBotID(botId string) error
}
type ComponentRepository interface {
	FindByFlowAndStepId(flowId string, stepId string) ([]Component, error)
	Create(component Component) (Component, error)
	Update(component Component) (Component, error)
	Delete(component Component) (Component, error)
	DeleteByBotID(botId string) error
}
