package templatepool

import "fmt"

func NewComponentService(repo ComponentRepository) ComponentService {
	return ComponentService{componentRepository: repo}
}

type ComponentService struct {
	componentRepository ComponentRepository
}

func (c ComponentService) NewComponent(component IComponent) (IComponent, error) {
	componentCreated, _ := c.componentRepository.Create(component)
	return componentCreated, nil
}

func (c ComponentService) UpdateComponent(component IComponent) (IComponent, error) {

	componentUpdated, _ := c.componentRepository.Update(component)
	return componentUpdated, nil
}

func (c ComponentService) DeleteComponent(component IComponent) (IComponent, error) {
	componentDeleted, _ := c.componentRepository.Delete(component)
	return componentDeleted, nil
}

func (c ComponentService) FindByFlowIdAndStepId(stepId string, flowId string) ([]IComponent, error) {
	components, _ := c.componentRepository.FindByFlowAndStepId(flowId, stepId)
	fmt.Println(components)
	return components, nil
}
