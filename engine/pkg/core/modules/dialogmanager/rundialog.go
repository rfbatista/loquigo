package dialogmanager

import (
	"loquigo/engine/pkg/core/domain"
)

func NewRunDialogService() RunDialogService {
	return RunDialogService{}
}

type RunDialogService struct{}

func (r RunDialogService) Run() ([]domain.Message, error) {
	return []domain.Message{}, nil
}
