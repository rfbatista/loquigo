package repositories

import (
	"context"
	"loquigo/engine/src/core/modules/template/pool"
	database "loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewFlowRepository(mongodb database.MongoDB) FlowRepository {
	flowcollection := mongodb.Collection("flow")
	return FlowRepository{collection: flowcollection}
}

type FlowRepository struct {
	collection mongo.Collection
}

func (f FlowRepository) FindByBotId(id string) ([]pool.Flow, error) {
	//Bot id should be ObjectId
	//ID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		primitive.E{Key: "bot_id", Value: id},
	}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	var schemas []schemas.FlowSchema
	cursor, err := f.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []pool.Flow{}, err
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &schemas); err != nil {
		return []pool.Flow{}, err
	}
	var components = []pool.Flow{}
	for _, schema := range schemas {
		components = append(components, schema.ToDomain())
	}
	return components, nil
}

func (f FlowRepository) Create(flow pool.Flow) (pool.Flow, error) {
	schema, _ := schemas.NewFlowSchema(flow)
	_, err := f.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return pool.Flow{}, err
	}
	return schema.ToDomain(), nil
}

func (f FlowRepository) Update(flow pool.Flow) (pool.Flow, error) {
	schema, _ := schemas.NewFlowSchema(flow)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{primitive.E{Key: "id", Value: schema.ID}}
	_, err := f.collection.UpdateOne(context.TODO(), filter, schema, opts)
	if err != nil {
		return pool.Flow{}, err
	}
	return schema.ToDomain(), nil
}

func (f FlowRepository) Delete(flow pool.Flow) (pool.Flow, error) {
	schema, _ := schemas.NewFlowSchema(flow)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "id", Value: schema.ID}}
	_, err := f.collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return pool.Flow{}, err
	}
	return schema.ToDomain(), nil
}

func (c FlowRepository) DeleteByBotID(botId string) error {
	opts := options.Delete()
	filter := bson.M{"bot_id": bson.M{"$eq": botId}}
	result, err := c.collection.DeleteMany(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {

	}
	if err != nil {
		return err
	}
	return nil
}
func (c FlowRepository) FindStepBeginIdFromFlow(flowId string) (string, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: flowId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.FlowSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return "", err
	}
	return schema.BeginId, nil
}
