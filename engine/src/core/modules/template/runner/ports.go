package runner

import (
	"loquigo/engine/src/core/domain"
)

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}

type ComponentRepository interface {
	FindByFlowAndStepId(flowId string, stepId string) ([]IComponent, error)
	Create(component IComponent) (IComponent, error)
	Update(component IComponent) (IComponent, error)
	Delete(component IComponent) (IComponent, error)
}
