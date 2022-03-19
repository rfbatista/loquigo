package components

import "loquigo/engine/pkg/core/domain"

func NewGotoComponent(c domain.Component) GotoComponent {
	return GotoComponent{domain.Component{Data: domain.ComponentData{GroupId: c.Data.GroupId, NodeId: c.Data.NodeId}}}
}

type GotoComponent struct {
	domain.Component
}

func (g GotoComponent) Run(m domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo) {
	return botMessages, nil, &domain.GoTo{GroupId: g.Data.GroupId, NodeId: g.Data.NodeId}
}
