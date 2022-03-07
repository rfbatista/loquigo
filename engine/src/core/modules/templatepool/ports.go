package templatepool

import (
	"loquigo/engine/src/core/domain"
)

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}

type FlowRepository interface {
	FindByBotId(id string) ([]Flow, error)
	Create(flow Flow) (Flow, error)
	Update(flow Flow) (Flow, error)
	Delete(flow Flow) (Flow, error)
}
type StepRepository interface {
	FindByFlowId(id string) ([]Step, error)
	FindById(id string) (Step, error)
	Create(step Step) (Step, error)
	Update(step Step) (Step, error)
	Delete(step Step) (Step, error)
}
type ComponentRepository interface {
	FindByFlowAndStepId(flowId string, stepId string) ([]IComponent, error)
	Create(component IComponent) (IComponent, error)
	Update(component IComponent) (IComponent, error)
	Delete(component IComponent) (IComponent, error)
}
