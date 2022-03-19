package schemas

import (
	"loquigo/engine/src/core/modules/components"
)

func NewComponentSchema(component components.Component) (ComponentSchema, error) {
	return ComponentSchema{
		ID:           component.ID,
		BotReference: component.BotReference,
		GroupId:      component.GroupId,
		NodeId:       component.NodeId,
		Type:         component.Type,
		Data:         component.Data,
		Sequence:     component.Sequence,
	}, nil
}

type ComponentSchema struct {
	ID           string                   `bson:"id" json:"id,omitempty"`
	BotReference string                   `bson:"bot_reference"`
	GroupId      string                   `bson:"group_id"`
	NodeId       string                   `bson:"group_id"`
	Type         string                   `bson:"type"`
	Data         components.ComponentData `bson:"data"`
	Sequence     int                      `bson:"sequence"`
}

func (c ComponentSchema) ToDomain() components.Component {
	return components.Component{
		ID:       c.ID,
		GroupId:  c.GroupId,
		NodeId:   c.NodeId,
		Type:     c.Type,
		Data:     c.Data,
		Sequence: c.Sequence,
	}
}
