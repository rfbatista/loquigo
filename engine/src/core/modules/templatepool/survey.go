package templatepool

import (
	"loquigo/engine/src/core/domain"
)

type SurveyInput struct {
	message domain.Message
	context domain.UserContext
}

type SurveyService struct {
}

func NewSurveyService() {}

func (s SurveyService) Run(i SurveyInput) []domain.Message {
	message, _, _ := domain.CreateText(string("hello world")).Run(i.message, i.context)
	return message
}
