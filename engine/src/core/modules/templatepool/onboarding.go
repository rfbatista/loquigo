package templatepool

import (
	"loquigo/engine/src/core/domain"
)

var (
	validatecode = "validateCode"
)

type Onboarding struct {
	name  string
	steps StepHash
}

func NewOnboarding() map[string]StepHash {
	var m = StepHash{
		"start":        Begin,
		"validateCode": ValidateCode,
	}
	flows := map[string]StepHash{
		"begin": m,
	}
	return flows
}

func Begin(userMessage domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component {
	var components = []domain.Component{
		domain.NewText("*Boas-vindas!* Para começar, por favor digite o código que você recebeu."),
		domain.NewHold("onboarding", "validateCode"),
	}
	return components
}

func ValidateCode(userMessage domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component {
	var components = []domain.Component{
		domain.NewText("Validando codigo..."),
		domain.NewHold("onboarding", "start"),
	}
	return components
}
