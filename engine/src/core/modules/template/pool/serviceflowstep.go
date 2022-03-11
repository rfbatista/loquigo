package pool

func NewStepService(repo StepRepository, service ComponentService) StepService {
	return StepService{stepRepo: repo, componentService: service}
}

type StepService struct {
	stepRepo         StepRepository
	componentService ComponentService
}

func (s StepService) NewStep(step Step) (Step, error) {

	stepCreated, _ := s.stepRepo.Create(step)
	return stepCreated, nil
}

func (s StepService) UpdateStep(step Step) (Step, error) {
	stepCreated, _ := s.stepRepo.Update(step)
	for _, component := range step.Components {
		s.componentService.UpdateComponent(component)
	}
	return stepCreated, nil
}

func (s StepService) DeleteStep(step Step) (Step, error) {
	stepCreated, _ := s.stepRepo.Delete(step)
	for _, component := range step.Components {
		s.componentService.UpdateComponent(component)
	}
	return stepCreated, nil
}

func (s StepService) FindByFlowId(flowId string) ([]Step, error) {
	steps, _ := s.stepRepo.FindByFlowId(flowId)

	var stepsWithComponents []Step
	for _, step := range steps {
		components, _ := s.componentService.FindByFlowIdAndStepId(step.ID, step.FlowId)
		step.Components = components
		stepsWithComponents = append(stepsWithComponents, step)
	}
	return stepsWithComponents, nil
}

func (s StepService) FindById(flowId string) (Step, error) {
	step, _ := s.stepRepo.FindById(flowId)
	components, _ := s.componentService.FindByFlowIdAndStepId(step.ID, step.FlowId)
	step.Components = components
	return step, nil
}

func (s StepService) DeleteByBotID(botId string) error {
	_ = s.stepRepo.DeleteByBotID(botId)
	return nil
}

func (s StepService) FindByIdFlowId(Id string, flowId string) (Step, error) {
	return Step{}, nil
}
