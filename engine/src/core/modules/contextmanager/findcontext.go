package contextmanager

import (
	"loquigo/engine/src/core/domain"
)

type FindContextService struct{}

func NewFindContextService() FindContextService {
	return FindContextService{}
}

func (f FindContextService) Run(event domain.Event) (domain.UserContext, error) {
	return domain.UserContext{}, nil
}
