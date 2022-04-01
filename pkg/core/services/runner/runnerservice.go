package runner

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/nodes"
)

func NewRunnerService(repo UserStateRepo, stepService nodes.RunnerNodeService, runner Runner) RunnerService {
	return RunnerService{userStateRepo: repo, runnerStepService: stepService, runner: runner}
}

type RunnerService struct {
	userStateRepo     UserStateRepo
	runnerStepService nodes.RunnerNodeService
	runner            Runner
}

func (t RunnerService) Run(event domain.Event, context domain.UserContext) ([]domain.Message, error) {
	state, err := t.FindUserState(event.Bot.ID, event.User)
	if err != nil {
		return []domain.Message{}, err
	}
	input := RunnerInput{
		User:         context.User,
		BotReference: event.Bot.ID,
		Message:      event.Message,
		Context:      context,
		State:        state,
	}
	messages, newState, err := t.runner.Run(input)
	if err != nil {
		return messages, err
	}
	t.userStateRepo.Update(domain.UserState{UserId: event.User.ID, State: domain.State{GroupId: newState.GroupId, NodeId: newState.NodeId}})
	return messages, nil
}

func (t RunnerService) FindUserState(botReference string, user domain.User) (domain.State, error) {
	state, _ := t.userStateRepo.FindByUserId(user.ID)
	if (domain.UserState{}) == state {
		groupId, err := t.runnerStepService.FindBotBegin(botReference)
		if err != nil {
			return domain.State{}, err
		}
		stepId, err := t.runnerStepService.FindGroupBeginId(botReference, groupId)
		if err != nil {
			return domain.State{}, err
		}
		if groupId == "" || stepId == "" {
			return domain.State{}, MissingValidBotState{BotReference: botReference}
		}
		return domain.State{GroupId: groupId, NodeId: stepId}, nil
	}
	return domain.State{GroupId: state.GroupId, NodeId: state.NodeId}, nil
}
