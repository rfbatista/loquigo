package repositories

import (
	"context"
	"loquigo/engine/pkg/core/domain"
	database "loquigo/engine/pkg/infrastructure/database/mongo"
	"loquigo/engine/pkg/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewGroupRepository(mongodb database.MongoDB) GroupRepository {
	flowcollection := mongodb.Collection("group")
	return GroupRepository{collection: flowcollection}
}

type GroupRepository struct {
	collection mongo.Collection
}

func (f GroupRepository) FindByBotId(id string) ([]domain.Group, error) {
	//Bot id should be ObjectId
	//ID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		primitive.E{Key: "bot_id", Value: id},
	}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	var schemas []schemas.GroupSchema
	cursor, err := f.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []domain.Group{}, err
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &schemas); err != nil {
		return []domain.Group{}, err
	}
	var components = []domain.Group{}
	for _, schema := range schemas {
		components = append(components, schema.ToDomain())
	}
	return components, nil
}

func (f GroupRepository) Create(group domain.Group) (domain.Group, error) {
	schema, _ := schemas.NewGroupSchema(group)
	_, err := f.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return domain.Group{}, err
	}
	return schema.ToDomain(), nil
}

func (f GroupRepository) Update(flow domain.Group) (domain.Group, error) {
	schema, _ := schemas.NewGroupSchema(flow)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{primitive.E{Key: "id", Value: schema.ID}}
	_, err := f.collection.UpdateOne(context.TODO(), filter, schema, opts)
	if err != nil {
		return domain.Group{}, err
	}
	return schema.ToDomain(), nil
}

func (f GroupRepository) Delete(group domain.Group) (domain.Group, error) {
	schema, _ := schemas.NewGroupSchema(group)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "id", Value: schema.ID}}
	_, err := f.collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return domain.Group{}, err
	}
	return schema.ToDomain(), nil
}

func (c GroupRepository) DeleteByBotID(botId string) error {
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
func (c GroupRepository) FindBeginId(botId string, groupId string) (string, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: groupId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.GroupSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return "", err
	}
	return schema.BeginId, nil
}
