package components

import "loquigo/engine/src/core/domain"

type RunnerComponent interface {
	Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo)
}

func BuildRunnerComponent(c domain.Component) RunnerComponent {
	switch c.Type {
	case "text":
		return domain.NewTextComponent(c)
	case "goto":
		return domain.NewGotoComponent(c)
	case "hold":
		return domain.NewHoldComponent(c)
	}
	//todo: need to avoid creating a component when there is no type matched
	return domain.NewTextComponent(c)
}
