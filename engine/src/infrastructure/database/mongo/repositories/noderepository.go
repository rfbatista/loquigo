package repositories

import (
	"context"
	"loquigo/engine/src/core/modules/nodes"
	database "loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewNodeRepository(mongodb database.MongoDB) NodeRepository {
	usersCollection := mongodb.Collection("flow_step")
	return NodeRepository{collection: usersCollection}
}

type NodeRepository struct {
	collection mongo.Collection
}

func (s NodeRepository) FindByGroupId(id string) ([]nodes.Node, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: id},
	}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	var schemas []schemas.NodeSchema
	cursor, err := s.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []nodes.Node{}, err
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &schemas); err != nil {
		return []nodes.Node{}, err
	}
	var components = []nodes.Node{}
	for _, schema := range schemas {
		components = append(components, schema.ToDomain())
	}
	return components, nil
}

func (s NodeRepository) FindById(id string) (nodes.Node, error) {
	ID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		primitive.E{Key: "_id", Value: ID},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schemas schemas.NodeSchema
	err := s.collection.FindOne(context.TODO(), filter, opts).Decode(&schemas)
	if err != nil {
		return nodes.Node{}, err
	}
	return schemas.ToDomain(), nil
}

func (s NodeRepository) Create(step nodes.Node) (nodes.Node, error) {
	schema, _ := schemas.NewNodeSchema(step)
	_, err := s.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return nodes.Node{}, err
	}
	return schema.ToDomain(), nil
}

func (s NodeRepository) Update(step nodes.Node) (nodes.Node, error) {
	schema, _ := schemas.NewNodeSchema(step)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{primitive.E{Key: "_id", Value: schema.ID}}
	_, err := s.collection.UpdateOne(context.TODO(), filter, schema, opts)
	if err != nil {
		return nodes.Node{}, err
	}
	return schema.ToDomain(), nil
}

func (s NodeRepository) Delete(step nodes.Node) (nodes.Node, error) {
	schema, _ := schemas.NewNodeSchema(step)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "id", Value: schema.ID}}
	_, err := s.collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return nodes.Node{}, err
	}
	return schema.ToDomain(), nil
}

func (s NodeRepository) DeleteByBotID(botId string) error {
	opts := options.Delete()
	filter := bson.M{"bot_id": bson.M{"$eq": botId}}
	result, err := s.collection.DeleteMany(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {

	}
	if err != nil {
		return err
	}
	return nil
}

func (c NodeRepository) FindByIdAndGroupId(groupId string, nodeId string) (nodes.Node, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: groupId},
		primitive.E{Key: "id", Value: nodeId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.NodeSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return nodes.Node{}, err
	}
	return schema.ToDomain(), nil
}

func (c NodeRepository) FindByGroupIdAndNodeId(botId string, groupId string, nodeId string) (nodes.Node, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: groupId},
		primitive.E{Key: "id", Value: nodeId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema []schemas.NodeSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return nodes.Node{}, err
	}
	return nodes.Node{}, nil
}
