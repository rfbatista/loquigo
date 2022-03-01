package templatepool

func NewStepService(repo StepRepository) StepService {
	return StepService{stepRepo: repo}
}

type StepService struct {
	stepRepo StepRepository
}

func (s StepService) NewStep(step Step) (Step, error) {
	stepCreated, _ := s.stepRepo.Create(step)
	return stepCreated, nil
}

func (s StepService) UpdateStep(step Step) (Step, error) {
	stepCreated, _ := s.stepRepo.Update(step)
	return stepCreated, nil
}

func (s StepService) DeleteStep(step Step) (Step, error) {
	stepCreated, _ := s.stepRepo.Delete(step)
	return stepCreated, nil
}

func (s StepService) FindByFlowId(flowId string) ([]Step, error) {
	steps, _ := s.stepRepo.FindByFlowId(flowId)
	return steps, nil
}
