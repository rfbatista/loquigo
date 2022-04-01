package repositories

import (
	"context"
	"log"
	"loquigo/engine/pkg/core/domain"
	database "loquigo/engine/pkg/infrastructure/database/mongo"
	"loquigo/engine/pkg/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewBotVersionRepository(mongo database.MongoDB) BotVersionRepository {
	collection := mongo.Collection("bot_version")
	return BotVersionRepository{collection: collection}
}

type BotVersionRepository struct {
	collection mongo.Collection
}

func (u BotVersionRepository) Create(bot domain.BotVersion) (domain.BotVersion, error) {
	botSchema := schemas.NewBotVersionSchema(bot)
	_, err := u.collection.InsertOne(context.Background(), botSchema)
	if err != nil {
		return bot, err
	}
	return botSchema.ToDomain(), nil
}

func (b BotVersionRepository) FindByBotId(botId string) ([]domain.BotVersion, error) {
	filter := bson.D{
		primitive.E{Key: "bot_id", Value: botId},
	}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	cur, err := b.collection.Find(context.Background(), filter, opts)
	defer cur.Close(context.Background())
	if err != nil {
		return []domain.BotVersion{}, err
	}
	var botVersions = []domain.BotVersion{}
	for cur.Next(context.TODO()) {
		var elem schemas.BotVersionSchema
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		botVersions = append(botVersions, elem.ToDomain())
	}
	return botVersions, nil
}

func (b BotVersionRepository) FindByIdAndBotId(versionId string, botId string) (domain.BotVersion, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: versionId},
		primitive.E{Key: "bot_id", Value: botId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.BotVersionSchema
	err := b.collection.FindOne(context.Background(), filter, opts).Decode(&schema)
	if err != nil {
		return domain.BotVersion{}, err
	}
	return schema.ToDomain(), nil
}
