package user

import (
	"context"
	"loquigo/engine/database"
	"loquigo/engine/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = "user"

type UserSchema struct {
	ID         primitive.ObjectID `bson:"_id"`
	ExternalId string             `bson:"external_id"`
	StateId    string             `bson:"state_id"`
	GroupId    string             `bson:"group_id"`
	BotId      string             `bson:"bot_id"`
}

func NewUserDAO(user *User) UserSchema {
	id := primitive.NewObjectID()
	return UserSchema{ID: id, ExternalId: user.ExternalId}
}

func (u *UserSchema) Save() {
	db := database.GetMongoConnection()
	db.Collection(collection).InsertOne(context.TODO(), u)
}

func toModel(schema UserSchema) *User {
	return &User{
		id: schema.ID.String(),
	}
}

func FindByExternalId(externalId string) (*User, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"external_id": externalId}
	var result UserSchema
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return toModel(result), nil
}

func FindById(id string) (*User, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result UserSchema
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return toModel(result), nil
}

func UpdateUserState(id string, state Location) error {
	db := database.GetMongoConnection()
	opts := options.Update().SetUpsert(false)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"state_id": state.StateID, "group_id": state.GroupID}}
	result, err := db.Collection(collection).UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		logger.Info("User state not updated")
		return nil
	}
	return nil
}
