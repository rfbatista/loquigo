package dialogmanager

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/runner"
)

func NewDialogManagerService(template runner.RunnerService, dialogmanager FindContextService) DialogManagerService {
	return DialogManagerService{runnerService: template, dialogmanager: dialogmanager}
}

type DialogManagerService struct {
	runnerService runner.RunnerService
	dialogmanager FindContextService
}

func (r DialogManagerService) Run(event domain.Event) ([]domain.Message, error) {
	userContext, _ := r.dialogmanager.Run(event)
	messages, err := r.runnerService.Run(event, userContext)
	if err != nil {
		return messages, err
	}
	return messages, nil
}