package repositories

import (
	"context"
	"loquigo/engine/src/core/domain"
	database "loquigo/engine/src/infrastructure/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewBotRepository(mongodb database.MongoDB) BotRepository {
	botCollection := mongodb.Collection("flow_component")
	return BotRepository{collection: botCollection}
}

type BotRepository struct {
	collection mongo.Collection
}

func (u BotRepository) Create(bot domain.Bot) (domain.Bot, error) {
	_, err := u.collection.InsertOne(context.TODO(), bot)
	if err != nil {
		return bot, err
	}
	return bot, nil
}

func (b BotRepository) FindBotBegin(botId string) (string, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: botId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema domain.Bot
	err := b.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return "", err
	}
	return schema.BeginId, nil
}
