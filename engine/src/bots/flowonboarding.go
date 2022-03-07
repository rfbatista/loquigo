package bots

// package templatepool

// import (
// 	"loquigo/engine/src/core/domain"
// )

// var (
// 	validatecode = "validateCode"
// )

// type Onboarding struct {
// 	name  string
// 	steps StepHash
// }

// func NewOnboarding() FlowHash {
// 	var m = StepHash{
// 		"start":        Begin,
// 		"validateCode": ValidateCode,
// 	}
// 	flows := FlowHash{
// 		"begin": m,
// 	}
// 	return flows
// }

// func Begin(userMessage domain.Message, context domain.UserContext, messages []domain.Message) []IComponent {
// 	var components = []IComponent{
// 		NewText("*Boas-vindas!* Para começar, por favor digite o código que você recebeu."),
// 		NewHold("begin", "validateCode"),
// 	}
// 	return components
// }

// func ValidateCode(userMessage domain.Message, context domain.UserContext, messages []domain.Message) []IComponent {
// 	var components = []IComponent{
// 		NewText("Validando codigo..."),
// 		NewHold("begin", "start"),
// 	}
// 	return components
// }
