package runner

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

func NewRunnerInput(message domain.Message, context domain.UserContext, state domain.State) RunnerInput {
	return RunnerInput{Message: message, Context: context, State: state}
}

type RunnerInput struct {
	User    domain.User
	BotId   string
	Message domain.Message
	Context domain.UserContext
	State   domain.State
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

	if (domain.State{}) == i.State {
		return botMessages, domain.State{}, MissingState{UserId: i.User.ID}
	}
	step, err := r.FindStep(i.BotId, i.State)
	if err != nil {
		return botMessages, i.State, err
	}
	if err != nil {
		return botMessages, domain.NewState(i.State.FlowId, i.State.StepId), err
	}

	for outer := 0; outer < 30; outer++ {
		botMessages, stop, goTo = step.Run(i.Message, i.Context, botMessages)
		if stop != nil {
			return botMessages, domain.NewState(stop.FlowId, stop.StepId), nil
		}
		if goTo == nil {
			return botMessages, domain.NewState(i.State.FlowId, i.State.StepId), nil
		}
		step, err = r.FindStep(i.BotId, domain.NewState(goTo.FlowId, goTo.StepId))
		if err != nil {
			return botMessages, domain.NewState(i.State.FlowId, i.State.StepId), err
		}
	}
	return botMessages, domain.NewState(i.State.FlowId, i.State.StepId), nil
}

func (r Runner) FindStep(botId string, state domain.State) (RunnerStep, error) {
	step, err := r.runnerStepService.FindByFlowIdAndStepId(botId, state.FlowId, state.StepId)
	if err != nil {
		return step, err
	}
	if step != nil {
		return step, nil
	}
	step, err = r.runnerStepService.FindFlowBegin(botId, state.FlowId)
	if err != nil {
		return step, err
	}
	if step != nil {
		return step, nil
	}
	flowId, err := r.runnerStepService.FindBotBegin(botId)
	if err != nil {
		return step, err
	}
	step, err = r.runnerStepService.FindFlowBegin(botId, flowId)
	if err != nil {
		return step, err
	}
	if step != nil {
		return step, nil
	}
	return step, NoStepInBotError{BotId: botId}
}
