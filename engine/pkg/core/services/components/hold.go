package components

import "loquigo/engine/pkg/core/domain"

func NewHoldComponent(c domain.Component) HoldComponent {
	return HoldComponent{Data: domain.ComponentData{GroupId: c.Data.GroupId, NodeId: c.Data.NodeId}}
}

type HoldComponent struct {
	Data domain.ComponentData
}

type HoldComponentData struct {
	GroupId string
	NodeId  string
}

func (h HoldComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	return botMessages, &domain.Stop{NodeId: h.Data.NodeId, GroupId: h.Data.GroupId}, nil
}
