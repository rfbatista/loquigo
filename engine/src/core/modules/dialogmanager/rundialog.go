package dialogmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/runner"
)

func NewRunDialogService(template runner.TemplateRunnerService, dialogmanager FindContextService) RunDialogService {
	return RunDialogService{template: template, dialogmanager: dialogmanager}
}

type RunDialogService struct {
	template      runner.TemplateRunnerService
	dialogmanager FindContextService
}

func (r RunDialogService) Run(event domain.Event) ([]domain.Message, error) {
	userContext, _ := r.dialogmanager.Run(event)
	messages, _ := r.template.Run(event, userContext)
	return messages, nil
}
