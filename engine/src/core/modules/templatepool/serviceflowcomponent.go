package templatepool

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

func (c ComponentService) GetComponents(stepId string, flowId string) ([]Component, error) {
	components, _ := c.componentRepository.FindByFlowAndStepId(flowId, stepId)
	return components, nil
}
