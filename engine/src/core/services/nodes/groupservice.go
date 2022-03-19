package nodes

import "loquigo/engine/src/core/domain"

func NewGroupService(repo GroupRepository) GroupService {
	return GroupService{groupRepository: repo}
}

type GroupService struct {
	groupRepository GroupRepository
}

func (f GroupService) NewGroup(flow domain.Group) (domain.Group, error) {
	flowCreated, _ := f.groupRepository.Create(flow)
	return flowCreated, nil
}

func (f GroupService) Update(flow domain.Group) (domain.Group, error) {
	flowUpdated, _ := f.groupRepository.Update(flow)
	return flowUpdated, nil
}

func (f GroupService) Delete(flow domain.Group) (domain.Group, error) {
	flowDeleted, _ := f.groupRepository.Delete(flow)
	return flowDeleted, nil
}

func (f GroupService) FindByBotId(botReference string) ([]domain.Group, error) {
	flows, _ := f.groupRepository.FindByBotId(botReference)
	return flows, nil
}

func (f GroupService) DeleteByBotID(botReference string) error {
	_ = f.groupRepository.DeleteByBotID(botReference)
	return nil
}
