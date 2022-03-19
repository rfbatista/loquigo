package components

type ComponentRepository interface {
	FindByGroupIdAndNodeId(botReference string, groupId string, nodeId string) ([]Component, error)
	Create(component Component) (Component, error)
	Update(component Component) (Component, error)
	Delete(component Component) (Component, error)
	DeleteByBotID(botReference string) error
}
