package templatepool

import (
	"loquigo/engine/src/core/domain"
)

func NewTemplatePoolService(repo UserStateRepo) TemplateRunnerService {
	return TemplateRunnerService{userStateRepo: repo}
}

type TemplateRunnerService struct {
	userStateRepo UserStateRepo
}

func (t TemplateRunnerService) Run(event domain.Event, context domain.UserContext) ([]domain.Message, error) {
	state, nil := t.userStateRepo.FindByUserId(event.User.ExternalId)
	input := NewRunnerInput(event.Message, context, state)
	onboarding := NewOnboarding()
	service := RunnerService{flow: onboarding}
	res, newState := service.Run(input)
	t.userStateRepo.Update(t.newUserState(state, newState))
	return res, nil
}

func (t TemplateRunnerService) newUserState(userState domain.UserState, state domain.State) domain.UserState {
	userState.FlowId = state.FlowId
	userState.StepId = state.StepId
	return userState
}
