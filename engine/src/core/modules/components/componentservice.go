package components

func NewComponentService(repo ComponentRepository) ComponentService {
	return ComponentService{componentRepository: repo}
}

type ComponentService struct {
	componentRepository ComponentRepository
}

func (c ComponentService) NewComponent(component Component) (Component, error) {
	componentCreated, _ := c.componentRepository.Create(component)
	return componentCreated, nil
}

func (c ComponentService) UpdateComponent(component Component) (Component, error) {
	componentUpdated, _ := c.componentRepository.Update(component)
	return componentUpdated, nil
}

func (c ComponentService) DeleteComponent(component Component) (Component, error) {
	componentDeleted, _ := c.componentRepository.Delete(component)
	return componentDeleted, nil
}

func (c ComponentService) FindByGroupIdAndNodeId(botReference string, nodeId string, groupId string) ([]Component, error) {
	components, _ := c.componentRepository.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	return components, nil
}

func (c ComponentService) DeleteByBotId(nodeID string) error {
	_ = c.componentRepository.DeleteByBotID(nodeID)
	return nil
}
