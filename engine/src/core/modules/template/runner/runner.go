package runner

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

func NewRunnerInput(message domain.Message, context domain.UserContext, state domain.State) RunnerInput {
	return RunnerInput{message: message, context: context, state: state}
}

type RunnerInput struct {
	user    domain.User
	botId   string
	message domain.Message
	context domain.UserContext
	state   domain.State
}

func NewRunnerService(stepService RunnerStepService) Runner {
	return Runner{runnerStepService: stepService}
}

type Runner struct {
	runnerStepService RunnerStepService
}

func (r Runner) Run(i RunnerInput) ([]domain.Message, domain.State, error) {
	var step RunnerStep
	var goTo *pool.GoTo
	var stop *pool.Stop
	var botMessages []domain.Message

	if (domain.State{}) == i.state {
		return botMessages, domain.State{}, MissingState{UserId: i.user.ID}
	}
	step, err := r.FindStep(i.botId, i.state)
	if err != nil {
		return botMessages, domain.NewState(i.state.FlowId, i.state.StepId), err
	}

	for outer := 0; outer < 30; outer++ {
		botMessages, stop, goTo = step.Run(i.message, i.context, botMessages)
		if stop != nil {
			return botMessages, domain.NewState(stop.FlowId, stop.StepId), nil
		}
		if goTo == nil {
			return botMessages, domain.NewState(i.state.FlowId, i.state.StepId), nil
		}
		step, err = r.FindStep(i.botId, domain.NewState(goTo.FlowId, goTo.StepId))
		if err != nil {
			return botMessages, domain.NewState(i.state.FlowId, i.state.StepId), err
		}
	}
	return botMessages, domain.NewState(i.state.FlowId, i.state.StepId), nil
}

func (r Runner) FindStep(botId string, state domain.State) (RunnerStep, error) {
	step, _ := r.runnerStepService.FindByFlowIdAndStepId(state.FlowId, state.StepId)
	if step != nil {
		return step, nil
	}
	step, _ = r.runnerStepService.FindFlowBegin(state.FlowId)
	if step != nil {
		return step, nil
	}
	flowId, _ := r.runnerStepService.FindBotBegin(botId)
	step, _ = r.runnerStepService.FindFlowBegin(flowId)
	if step != nil {
		return step, nil
	}
	return step, NoStepInBotError{BotId: botId}
}
