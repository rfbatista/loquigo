package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewComponentSchema(component pool.Component) (ComponentSchema, error) {
	ID, IDError := primitive.ObjectIDFromHex(component.ID)
	if IDError != nil {
		ID = primitive.NewObjectID()
	}
	return ComponentSchema{
		ID:       ID,
		Flow_id:  component.FlowId,
		Step_id:  component.StepId,
		Type:     component.Type,
		Data:     component.Data,
		Sequence: component.Sequence,
	}, nil
}

type ComponentSchema struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Flow_id  string             `bson:"flow_id"`
	Step_id  string             `bson:"step_id"`
	Type     string             `bson:"type"`
	Data     pool.ComponentData `bson:"data"`
	Sequence int                `bson:"sequence"`
}

func (c ComponentSchema) ToDomain() pool.Component {
	return pool.Component{
		ID:       c.ID.Hex(),
		FlowId:   c.Flow_id,
		StepId:   c.Step_id,
		Type:     c.Type,
		Data:     c.Data,
		Sequence: c.Sequence,
	}
}
