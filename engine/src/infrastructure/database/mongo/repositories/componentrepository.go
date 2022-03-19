package repositories

import (
	"context"
	"log"
	"loquigo/engine/src/core/modules/components"
	database "loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewComponentRepo(mongodb database.MongoDB) ComponentRepository {
	usersCollection := mongodb.Collection("component")
	return ComponentRepository{collection: usersCollection}
}

type ComponentRepository struct {
	collection mongo.Collection
}

func (c ComponentRepository) FindByGroupIdAndNodeId(botId string, groupId string, nodeId string) ([]components.Component, error) {
	filter := bson.M{"bot_id": botId, "group_id": groupId, "node_id": nodeId}
	opts := options.Find()
	cur, err := c.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []components.Component{}, err
	}
	defer cur.Close(context.TODO())
	var components = []components.Component{}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem schemas.ComponentSchema
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		components = append(components, elem.ToDomain())
	}
	return components, nil
}

func (c ComponentRepository) Create(component components.Component) (components.Component, error) {
	schema, _ := schemas.NewComponentSchema(component)
	_, err := c.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return components.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) Update(component components.Component) (components.Component, error) {
	schema, _ := schemas.NewComponentSchema(component)

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": schema.ID}
	//todo: remove direct fields insertion
	update := bson.D{{"$set", bson.M{
		"bot_reference": schema.BotReference,
		"group_id":      schema.GroupId,
		"node_id":       schema.NodeId,
		"type":          schema.Type,
		"data":          schema.Data,
		"sequence":      schema.Sequence},
	}}
	_, err := c.collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {

		return components.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) Delete(Icomponent components.Component) (components.Component, error) {
	schema, _ := schemas.NewComponentSchema(Icomponent)
	opts := options.Delete()
	filter := bson.M{"_id": schema.ID}
	result, err := c.collection.DeleteOne(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {

	}
	if err != nil {
		return components.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) DeleteByBotID(botId string) error {
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
