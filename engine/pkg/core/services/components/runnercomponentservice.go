package components

func NewRunnerComponentService(repo ComponentRepository) RunnerComponentService {
	return RunnerComponentService{componentRepository: repo}
}

type RunnerComponentService struct {
	componentRepository ComponentRepository
}

func (c RunnerComponentService) FindByGroupIdAndNodeId(botReference string, nodeId string, groupId string) ([]RunnerComponent, error) {
	var runnerComponents []RunnerComponent
	components, _ := c.componentRepository.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	for _, component := range components {
		runnerComponent := BuildRunnerComponent(component)
		runnerComponents = append(runnerComponents, runnerComponent)
	}
	return runnerComponents, nil
}
