package templatepool

import (
	"loquigo/engine/pkg/core/domain"
)

type Input struct {
	message domain.Message
	context domain.UserContext
}

type SurveyService struct {
}

func NewSurveyService() {}

func (s SurveyService) Run(i Input) []domain.Message {
	t := domain.TextComponent{}
	message, _, _ := t.Run(i.message, i.context)
	return message
}
