package repositories

import (
	"context"
	"fmt"
	"log"
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/infrastructure"
	database "loquigo/engine/pkg/infrastructure/database/mongo"
	"loquigo/engine/pkg/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(mongodb database.MongoDB) UserRepository {
	usersCollection := mongodb.Collection("user")
	return UserRepository{client: mongodb, collection: usersCollection}
}

type UserRepository struct {
	client     database.MongoDB
	collection mongo.Collection
	logger     infrastructure.Logger
}

func (u UserRepository) FindUserOrCreate(userExternalId string) (domain.User, error) {
	var user domain.User
	error := u.collection.FindOne(context.TODO(), bson.M{"external_id": userExternalId}).Decode(&user)
	if error != nil {
		fmt.Printf("Error: \n %+v", error)
		userToCreate := domain.User{ExternalId: userExternalId}
		userCreated, error := u.Create(userToCreate)
		if error != nil {
			return user, error
		}
		return userCreated, nil
	}
	return user, nil
}

func (u UserRepository) Create(user domain.User) (domain.User, error) {
	log.Println("Criando usuario")
	_, err := u.collection.InsertOne(context.TODO(), schemas.NewUserSchema(user))
	if err != nil {
		u.logger.Error("UserRepo", err)
		return user, err
	}
	return user, nil
}
