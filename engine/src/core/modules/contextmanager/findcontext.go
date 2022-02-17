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
	return domain.UserContext{}, nil
}
