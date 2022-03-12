package runner

import (
	"loquigo/engine/src/core/domain"
)

func NewChatRunnerService(repo UserStateRepo, stepService RunnerStepService, runner Runner) RunnerService {
	return RunnerService{userStateRepo: repo, runnerStepService: stepService, runner: runner}
}

type RunnerService struct {
	userStateRepo     UserStateRepo
	runnerStepService RunnerStepService
	runner            Runner
}

func (t RunnerService) Run(event domain.Event, context domain.UserContext) ([]domain.Message, error) {
	state, _ := t.FindUserState(event.Bot.ID, event.User)
	input := RunnerInput{
		User:    context.User,
		BotId:   event.Bot.ID,
		Message: event.Message,
		Context: context,
		State:   state,
	}
	messages, newState, err := t.runner.Run(input)
	if err != nil {
		return messages, err
	}

	t.userStateRepo.Update(domain.UserState{UserId: event.User.ID, FlowId: newState.FlowId, StepId: newState.StepId})
	return messages, nil
}

func (t RunnerService) FindUserState(botId string, user domain.User) (domain.State, error) {
	state, _ := t.userStateRepo.FindByUserId(user.ID)
	if (domain.UserState{}) == state {
		flowId, _ := t.runnerStepService.FindBotBegin(botId)
		stepId, _ := t.runnerStepService.FindFlowBeginId(flowId)
		return domain.State{FlowId: flowId, StepId: stepId}, nil
	}
	return domain.State{FlowId: state.FlowId, StepId: state.StepId}, nil
}
