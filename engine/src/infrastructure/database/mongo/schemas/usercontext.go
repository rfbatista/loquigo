package schemas

import (
	"loquigo/engine/src/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUserContextSchema(user domain.UserContext) UserContextSchema {
	var id primitive.ObjectID
	if user.ID == "" {
		id = primitive.NewObjectID()
	} else {
		id, _ = primitive.ObjectIDFromHex(user.ID)
	}
	return UserContextSchema{ID: id, UserId: user.UserId, BotId: user.BotId, User: user.User, Memory: user.Memory}
}

type UserContextSchema struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserId string             `bson:"user_id"`
	BotId  string             `bson:"bot_id"`
	User   domain.User        `bson:"user"`
	Memory map[string]string  `bson:"memory"`
}

func (u UserContextSchema) ToDomain() domain.UserContext {
	return domain.UserContext{ID: u.ID.Hex(), UserId: u.UserId, BotId: u.BotId, User: u.User, Memory: u.Memory}
}
