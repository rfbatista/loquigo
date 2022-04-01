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

func NewBotRepository(mongodb database.MongoDB) BotRepository {
	botCollection := mongodb.Collection("bot")
	botCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}, {Key: "version", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return BotRepository{collection: botCollection}
}

type BotRepository struct {
	collection mongo.Collection
}

func (u BotRepository) Create(bot domain.Bot) (domain.Bot, error) {
	_, err := u.collection.InsertOne(context.Background(), bot)
	if err != nil {
		return bot, err
	}
	return bot, nil
}

func (u BotRepository) Update(bot domain.Bot) (domain.Bot, error) {
	filter := bson.D{{"id", bot.ID}}
	update := bson.D{{"$set", bson.D{{"version", bot.CurrentVersion}}}}
	opts := options.Update().SetUpsert(false)
	// TODO: talvez devo montar o schema e retornar toda a autalizacao
	_, err := u.collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return bot, err
	}
	return bot, nil
}

func (b BotRepository) FindById(botId string) (domain.Bot, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: botId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.BotSchema
	err := b.collection.FindOne(context.Background(), filter, opts).Decode(&schema)
	if err != nil {
		return domain.Bot{}, err
	}
	return schema.ToDomain(), nil
}

func (b BotRepository) FindBeginByBotId(botId string) (string, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: botId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema domain.Bot
	err := b.collection.FindOne(context.Background(), filter, opts).Decode(&schema)
	if err != nil {
		return "", err
	}
	return schema.BeginId, nil
}

func (b BotRepository) GetBots() ([]domain.Bot, error) {
	cursor, err := b.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []domain.Bot{}, err
	}
	var botSchemas []schemas.BotSchema
	if err = cursor.All(context.Background(), &botSchemas); err != nil {
		log.Fatal(err)
	}
	var bots []domain.Bot
	for _, schema := range botSchemas {
		bots = append(bots, schema.ToDomain())
	}
	return bots, nil
}

func (b BotRepository) Delete(bot domain.Bot) (domain.Bot, error) {
	schema := schemas.NewBotSchema(bot)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "id", Value: schema.Id}}
	_, err := b.collection.DeleteOne(context.Background(), filter, opts)
	if err != nil {
		return domain.Bot{}, err
	}
	return schema.ToDomain(), nil
}
