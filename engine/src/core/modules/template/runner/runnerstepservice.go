package runner

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type RunnerStep interface {
	Run(message domain.Message, context domain.UserContext, messages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo)
	AddComponents(components []RunnerComponent) (RunnerStep, error)
	IsValid() bool
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

func (r RunnerStepService) FindFlowBegin(botId string, flowId string) (RunnerStep, error) {
	stepId, _ := r.flow.FindStepBeginIdFromFlow(botId, flowId)
	step, _ := r.stepRepo.FindByFlowIdAndStepId(botId, flowId, stepId)
	components, _ := r.componentRepo.FindByFlowIdAndStepId(botId, flowId, stepId)
	stepWithComponents, _ := step.AddComponents(components)
	if stepWithComponents.IsValid() == false {
		return stepWithComponents, InvalidStep{BotId: botId, FlowId: flowId, StepId: stepId}
	}
	return stepWithComponents, nil
}

func (r RunnerStepService) FindByFlowIdAndStepId(botId string, flowId string, stepId string) (RunnerStep, error) {
	step, _ := r.stepRepo.FindByFlowIdAndStepId(botId, flowId, stepId)
	components, _ := r.componentRepo.FindByFlowIdAndStepId(botId, flowId, stepId)
	stepWithComponents, _ := step.AddComponents(components)
	if stepWithComponents.IsValid() == false {
		return stepWithComponents, InvalidStep{BotId: botId, FlowId: flowId, StepId: stepId}
	}
	return stepWithComponents, nil
}
