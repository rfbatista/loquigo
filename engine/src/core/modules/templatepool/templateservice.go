package templatepool

import (
	"log"
	"loquigo/engine/src/core/domain"
)

func NewTemplatePoolService() TemplatePoolService {
	return TemplatePoolService{}
}

type TemplatePoolService struct{}

func (t TemplatePoolService) Run(message domain.Message, context domain.UserContext) ([]domain.Message, error) {
	input := RunnerInput{message: message, context: context}
	onboarding := NewOnboarding()
	service := RunnerService{flow: onboarding}
	res, _ := service.Run(input)
	log.Println(res)
	return res, nil
}
