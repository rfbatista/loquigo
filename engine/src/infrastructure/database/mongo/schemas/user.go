package schemas

import (
	"loquigo/engine/src/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUserSchema(user domain.User) UserSchema {
	id := primitive.NewObjectID()
	return UserSchema{ID: id, ExternalId: user.ExternalId}
}

type UserSchema struct {
	ID         primitive.ObjectID `bson:"_id"`
	ExternalId string             `bson:"external_id"`
}
