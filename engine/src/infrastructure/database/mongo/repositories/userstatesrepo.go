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

func NewUserStatestRepo(mongodb database.MongoDB) UserStatesRepo {
	usersCollection := mongodb.Collection("user_states")
	return UserStatesRepo{client: mongodb, collection: usersCollection}
}

type UserStatesRepo struct {
	client     database.MongoDB
	collection mongo.Collection
	logger     infrastructure.Logger
}

func (u UserStatesRepo) FindByUserId(id string) (*domain.UserState, error) {
	filter := bson.D{primitive.E{Key: "user_id", Value: id}}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var userState domain.UserState
	error := u.collection.FindOne(context.TODO(), filter, opts).Decode(&userState)
	if error != nil {
		u.logger.Error("UserStatesRepo", error)
		return nil, error
	}
	return &userState, nil
}

func (u UserStatesRepo) Update(userState domain.UserState) error {
	filter := bson.D{primitive.E{Key: "_id", Value: userState.ID}}
	_, err := u.collection.ReplaceOne(context.Background(), filter, userState)
	if err != nil {
		u.logger.Error("UserStatesRepo", err)
		return err
	}
	return nil
}

func (u UserStatesRepo) Create(userState domain.UserState) error {
	_, err := u.collection.InsertOne(context.TODO(), userState)
	if err != nil {
		u.logger.Error("UserStatesRepo", err)
		return err
	}
	return nil
}
