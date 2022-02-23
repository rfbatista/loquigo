package contextmanager

import (
	"loquigo/engine/src/core/domain"
)

type FindContextService struct {
	repo UserContextRepository
}

func NewFindContextService(repo UserContextRepository) FindContextService {
	return FindContextService{repo: repo}
}

func (f FindContextService) Run(event domain.Event) (domain.UserContext, error) {
	//todo: need to add bot in in this repo
	//context, err := f.repo.FindByUserId(event.User.ID)
	//if err != err {
	//	return domain.UserContext{}, err
	//}
	return domain.NewUserContext(), nil
}
