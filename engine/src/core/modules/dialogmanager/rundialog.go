package dialogmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/contextmanager"
	"loquigo/engine/src/core/modules/templatepool"
)

func NewRunDialogService() RunDialogService {
	return RunDialogService{}
}

type RunDialogService struct {
	template       templatepool.TemplatePoolService
	contextManager contextmanager.FindContextService
}

func (r RunDialogService) Run(event domain.Event) ([]domain.Message, error) {
	userContext, _ := r.contextManager.Run(event)
	messages, _ := r.template.Run(event.Message, userContext)
	return messages, nil
}
