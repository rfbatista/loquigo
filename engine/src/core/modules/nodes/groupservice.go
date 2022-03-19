package nodes

func NewGroupService(repo GroupRepository) GroupService {
	return GroupService{groupRepository: repo}
}

type GroupService struct {
	groupRepository GroupRepository
}

func (f GroupService) NewGroup(flow Group) (Group, error) {
	flowCreated, _ := f.groupRepository.Create(flow)
	return flowCreated, nil
}

func (f GroupService) Update(flow Group) (Group, error) {
	flowUpdated, _ := f.groupRepository.Update(flow)
	return flowUpdated, nil
}

func (f GroupService) Delete(flow Group) (Group, error) {
	flowDeleted, _ := f.groupRepository.Delete(flow)
	return flowDeleted, nil
}

func (f GroupService) FindByBotId(botReference string) ([]Group, error) {
	flows, _ := f.groupRepository.FindByBotId(botReference)
	return flows, nil
}

func (f GroupService) DeleteByBotID(botReference string) error {
	_ = f.groupRepository.DeleteByBotID(botReference)
	return nil
}
