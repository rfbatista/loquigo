package templatepool

import (
	"loquigo/engine/src/core/domain"
)

func NewRunnerInput(message domain.Message, context domain.UserContext, state domain.UserState) RunnerInput {
	return RunnerInput{message: message, context: context, state: state}
}

type RunnerInput struct {
	message domain.Message
	context domain.UserContext
	state   domain.UserState
}

func NewRunnerService(flows FlowHash) RunnerService {
	return RunnerService{flow: flows}
}

type RunnerService struct {
	flow FlowHash
}

func (s RunnerService) Run(i RunnerInput) ([]domain.Message, domain.State) {
	if i.state == (domain.UserState{}) {
		i.state = domain.NewUserState("", "begin", "start")
	}
	var flowId string = i.state.FlowId
	var stepId string = i.state.StepId
	var goTo *domain.GoTo
	var stop *domain.Stop
	messages := []domain.Message{}
	step := s.flow[flowId][stepId]
	for outer := 0; outer < 30; outer++ {
		messages, stop, goTo = s.RunFlow(step, i, messages)
		if stop != nil {
			return messages, domain.NewState(stop.FlowId, stop.StepId)
		}
		if goTo == nil {
			break
		}
	}
	return messages, domain.NewState(i.state.FlowId, i.state.StepId)
}

func (s RunnerService) RunFlow(startStep Step, i RunnerInput, previous []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	currentStep := startStep
	var changeFlow *domain.GoTo
	var stopFlow *domain.Stop
	var previousMessages []domain.Message = previous
	if startStep == nil {
		return []domain.Message{}, &domain.Stop{FlowId: i.state.FlowId, StepId: i.state.StepId}, &domain.GoTo{}
	}
	for circuitBreaker := 0; circuitBreaker < 30; circuitBreaker++ {
		previousMessages, stopFlow, changeFlow = s.RunStep(currentStep, i, previousMessages)
		if stopFlow != nil {
			return previousMessages, stopFlow, changeFlow
		}
		if changeFlow == nil {
			return previousMessages, stopFlow, changeFlow
		}
		currentStep = s.flow[changeFlow.FlowId][changeFlow.StepId]
	}
	return previousMessages, stopFlow, changeFlow
}

func (s RunnerService) RunStep(step Step, i RunnerInput, messages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	if step == nil {
		return []domain.Message{}, nil, nil
	}
	var changeFlow *domain.GoTo
	var stopFlow *domain.Stop
	var botMessages []domain.Message = messages
	for _, component := range step(i.message, i.context, messages) {
		botMessages, stopFlow, changeFlow = component.Run(i.message, i.context, botMessages)
	}
	return botMessages, stopFlow, changeFlow
}
