package schemas

import (
	"loquigo/engine/src/core/modules/templatepool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFlowSchema(flow templatepool.Flow) (FlowSchema, error) {
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

func (f FlowSchema) ToDomain() templatepool.Flow {
	return templatepool.NewFlow(f.ID.Hex(), f.BotId, f.Name)
}
