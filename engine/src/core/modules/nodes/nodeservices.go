package nodes

import "loquigo/engine/src/core/modules/components"

func NewNodeService(repo NodeRepository, service components.ComponentService) NodeService {
	return NodeService{nodeRepository: repo, componentService: service}
}

type NodeService struct {
	nodeRepository   NodeRepository
	componentService components.ComponentService
}

func (s NodeService) NewNode(step Node) (Node, error) {
	stepCreated, _ := s.nodeRepository.Create(step)
	return stepCreated, nil
}

func (s NodeService) UpdateNode(step Node) (Node, error) {
	stepCreated, _ := s.nodeRepository.Update(step)
	return stepCreated, nil
}

func (s NodeService) DeleteNode(step Node) (Node, error) {
	stepCreated, _ := s.nodeRepository.Delete(step)
	return stepCreated, nil
}

func (s NodeService) FindByGroupId(groupId string) ([]Node, error) {
	steps, _ := s.nodeRepository.FindByGroupId(groupId)
	return steps, nil
}

func (s NodeService) FindById(groupId string) (Node, error) {
	step, _ := s.nodeRepository.FindById(groupId)
	return step, nil
}

func (s NodeService) DeleteByBotID(botReference string) error {
	_ = s.nodeRepository.DeleteByBotID(botReference)
	return nil
}

//todo: implement
func (s NodeService) FindByIdGroupId(Id string, groupId string) (Node, error) {
	return Node{}, nil
}
