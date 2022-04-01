package runner

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/nodes"
)

func NewRunnerInput(message domain.Message, context domain.UserContext, state domain.State) RunnerInput {
	return RunnerInput{Message: message, Context: context, State: state}
}

type RunnerInput struct {
	User         domain.User
	BotReference string
	Message      domain.Message
	Context      domain.UserContext
	State        domain.State
}

func NewRunner(stepService nodes.RunnerNodeService) Runner {
	return Runner{runnerStepService: stepService}
}

type Runner struct {
	runnerStepService nodes.RunnerNodeService
}

func (r Runner) Run(i RunnerInput) ([]domain.Message, domain.State, error) {
	var step nodes.RunnerNode
	var goTo *domain.GoTo
	var stop *domain.Stop
	var botMessages []domain.Message

	if (domain.State{}) == i.State {
		return botMessages, domain.State{}, MissingUserState{UserId: i.User.ID}
	}
	step, err := r.FindStep(i.BotReference, i.State)
	if err != nil {
		return botMessages, i.State, err
	}
	if err != nil {
		return botMessages, domain.NewState(i.State.GroupId, i.State.NodeId), err
	}

	for outer := 0; outer < 30; outer++ {
		botMessages, stop, goTo = step.Run(i.Message, i.Context, botMessages)
		if stop != nil {
			return botMessages, domain.NewState(stop.GroupId, stop.NodeId), nil
		}
		if goTo == nil {
			return botMessages, domain.NewState(i.State.GroupId, i.State.NodeId), nil
		}
		step, err = r.FindStep(i.BotReference, domain.NewState(goTo.GroupId, goTo.NodeId))
		if err != nil {
			return botMessages, domain.NewState(i.State.GroupId, i.State.NodeId), err
		}
	}
	return botMessages, domain.NewState(i.State.GroupId, i.State.NodeId), nil
}

func (r Runner) FindStep(botReference string, state domain.State) (nodes.RunnerNode, error) {
	node, err := r.runnerStepService.FindByGroupIdAndNodeId(botReference, state.GroupId, state.NodeId)
	if err != nil {
		return node, err
	}
	if node != nil {
		return node, nil
	}
	node, err = r.runnerStepService.FindGroupBegin(botReference, state.GroupId)
	if err != nil {
		return node, err
	}
	if node != nil {
		return node, nil
	}
	groupId, err := r.runnerStepService.FindBotBegin(botReference)
	if err != nil {
		return node, err
	}
	node, err = r.runnerStepService.FindGroupBegin(botReference, groupId)
	if err != nil {
		return node, err
	}
	if node != nil {
		return node, nil
	}
	return node, NoStepInBotError{BotReference: botReference}
}
