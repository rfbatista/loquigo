package runner

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type RunnerStep interface {
	Run(message domain.Message, context domain.UserContext, messages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo)
	AddComponents(components []RunnerComponent) (RunnerStep, error)
}

func NewRunnerStepService(bot BotRepository, flow FlowRepository, stepRepo StepRepository, c ComponentRepository) RunnerStepService {
	return RunnerStepService{bot: bot, flow: flow, stepRepo: stepRepo, componentRepo: c}
}

type RunnerStepService struct {
	bot           BotRepository
	flow          FlowRepository
	stepRepo      StepRepository
	componentRepo ComponentRepository
}

func (r RunnerStepService) FindBotBegin(botId string) (string, error) {
	flow, _ := r.bot.FindBotBegin(botId)
	return flow, nil
}

func (r RunnerStepService) FindFlowBeginId(flowId string) (string, error) {
	stepId, _ := r.FindFlowBeginId(flowId)
	return stepId, nil
}

func (r RunnerStepService) FindFlowBegin(flowId string) (RunnerStep, error) {
	stepId, _ := r.flow.FindStepBeginIdFromFlow(flowId)
	step, _ := r.stepRepo.FindByFlowIdAndStepId(flowId, stepId)
	components, _ := r.componentRepo.FindByFlowIdAndStepId(flowId, stepId)
	stepWithComponents, _ := step.AddComponents(components)
	return stepWithComponents, nil
}

func (r RunnerStepService) FindByFlowIdAndStepId(flowId string, stepId string) (RunnerStep, error) {
	step, _ := r.stepRepo.FindByFlowIdAndStepId(flowId, stepId)
	components, _ := r.componentRepo.FindByFlowIdAndStepId(flowId, stepId)
	stepWithComponents, _ := step.AddComponents(components)

	return stepWithComponents, nil
}
