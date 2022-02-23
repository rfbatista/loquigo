package dialogmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/contextmanager"
	"loquigo/engine/src/core/modules/templatepool"
)

func NewRunDialogService(template templatepool.TemplatePoolService, contextManager contextmanager.FindContextService) RunDialogService {
	return RunDialogService{template: template, contextManager: contextManager}
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
