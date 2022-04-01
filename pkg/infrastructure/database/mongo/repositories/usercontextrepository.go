package repositories

import (
	"context"
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/infrastructure"
	database "loquigo/engine/pkg/infrastructure/database/mongo"
	"loquigo/engine/pkg/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserContextRepo(mongodb database.MongoDB) UserContextRepository {
	usersCollection := mongodb.Collection("user_context")
	return UserContextRepository{client: mongodb, collection: usersCollection}
}

type UserContextRepository struct {
	client     database.MongoDB
	collection mongo.Collection
	logger     infrastructure.Logger
}

func (u UserContextRepository) FindByUserId(userId string) (domain.UserContext, error) {
	filter := bson.D{primitive.E{Key: "user_id", Value: userId}}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var userContext schemas.UserContextSchema
	error := u.collection.FindOne(context.TODO(), filter, opts).Decode(&userContext)
	if error != nil {
		u.logger.Error("UserContextRepo", error)
		return domain.NewUserContext(), error
	}
	return userContext.ToDomain(), nil
}

func (u UserContextRepository) SaveMemory(ctx domain.UserContext) error {
	return nil
}

func (u UserContextRepository) Update(ctx domain.UserContext) error {
	schemas.NewUserContextSchema(ctx)
	filter := bson.D{primitive.E{Key: "_id", Value: ctx.ID}}
	_, err := u.collection.ReplaceOne(context.Background(), filter, ctx)
	if err != nil {
		u.logger.Error("UserContextRepo", err)
		return err
	}
	return nil
}

func (u UserContextRepository) Create(userContext domain.UserContext) {
	_, err := u.collection.InsertOne(context.TODO(), userContext)
	if err != nil {
		u.logger.Error("UserContextRepo", err)
	}
}
