package dialogmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/runner"
)

func NewRunDialogService(template runner.RunnerService, dialogmanager FindContextService) RunDialogService {
	return RunDialogService{runnerService: template, dialogmanager: dialogmanager}
}

type RunDialogService struct {
	runnerService runner.RunnerService
	dialogmanager FindContextService
}

func (r RunDialogService) Run(event domain.Event) ([]domain.Message, error) {
	userContext, _ := r.dialogmanager.Run(event)
	messages, _ := r.runnerService.Run(event, userContext)
	return messages, nil
}
