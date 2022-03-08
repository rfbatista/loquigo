package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFlowSchema(flow pool.Flow) (FlowSchema, error) {
	ID, _ := primitive.ObjectIDFromHex(flow.ID)
	return FlowSchema{
		ID:    ID,
		BotId: flow.BotId,
		Name:  flow.Name,
	}, nil
}

type FlowSchema struct {
	ID    primitive.ObjectID `bson:"_id"`
	BotId string             `bson:"bot_id"`
	Name  string             `bson:"name"`
}

func (f FlowSchema) ToDomain() pool.Flow {
	return pool.NewFlow(f.ID.Hex(), f.BotId, f.Name)
}
