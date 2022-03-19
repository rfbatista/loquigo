package components

import "loquigo/engine/src/core/domain"

type ComponentRepository interface {
	FindByGroupIdAndNodeId(botReference string, groupId string, nodeId string) ([]domain.Component, error)
	Create(component domain.Component) (domain.Component, error)
	Update(component domain.Component) (domain.Component, error)
	Delete(component domain.Component) (domain.Component, error)
	DeleteByBotID(botReference string) error
}
