package templatepool

import "loquigo/engine/src/core/domain"

func NewTemplatePoolService() TemplatePoolService {
	return TemplatePoolService{}
}

type TemplatePoolService struct{}

func (t TemplatePoolService) Run(message domain.Message, context domain.UserContext) ([]domain.Message, error) {
	input := SurveyInput{message: message, context: context}
	res := SurveyService{}.Run(input)
	return res, nil
}
