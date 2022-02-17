package repositories

import (
	"context"
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/infrastructure"
	database "loquigo/engine/src/infrastructure/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserContextRepo(mongodb database.MongoDB, logger infrastructure.Logger, cfg infrastructure.Config) UserContextRepo {
	usersCollection := mongodb.Client.Database(cfg.DbName()).Collection("user_context")
	return UserContextRepo{client: mongodb, collection: *usersCollection, logger: logger}
}

type UserContextRepo struct {
	client     database.MongoDB
	collection mongo.Collection
	logger     infrastructure.Logger
}

func (u UserContextRepo) FindByUserId(id string) (*domain.UserContext, error) {
	filter := bson.D{primitive.E{Key: "user_id", Value: id}}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var userContext domain.UserContext
	error := u.collection.FindOne(context.TODO(), filter, opts).Decode(&userContext)
	if error != nil {
		u.logger.Error("UserContextRepo", error)
		return nil, error
	}
	return &userContext, nil
}

func (u UserContextRepo) Update(userContext domain.UserContext) error {
	filter := bson.D{primitive.E{Key: "_id", Value: userContext.ID}}
	_, err := u.collection.ReplaceOne(context.Background(), filter, userContext)
	if err != nil {
		u.logger.Error("UserContextRepo", err)
		return err
	}
	return nil
}

func (u UserContextRepo) Create(userContext domain.UserContext) {
	_, err := u.collection.InsertOne(context.TODO(), userContext)
	if err != nil {
		u.logger.Error("UserContextRepo", err)
	}
}
