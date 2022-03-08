package pool

func NewFlowService(repo FlowRepository) FlowService {
	return FlowService{flowRepository: repo}
}

type FlowService struct {
	flowRepository FlowRepository
}

func (f FlowService) NewFlow(flow Flow) (Flow, error) {
	flowCreated, _ := f.flowRepository.Create(flow)
	return flowCreated, nil
}

func (f FlowService) Update(flow Flow) (Flow, error) {
	flowUpdated, _ := f.flowRepository.Update(flow)
	return flowUpdated, nil
}

func (f FlowService) Delete(flow Flow) (Flow, error) {
	flowDeleted, _ := f.flowRepository.Delete(flow)
	return flowDeleted, nil
}

func (f FlowService) FindByBotId(botId string) ([]Flow, error) {
	flows, _ := f.flowRepository.FindByBotId(botId)
	return flows, nil
}
