package mongo

import (
	"context"
	"log"
	"loquigo/engine/src/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var ctx = context.TODO()

func NewMongoDb(cfg infrastructure.Config) MongoDB {
	return MongoDB{}
}

type MongoDB struct {
	Client mongo.Client
}

func (m *MongoDB) Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	m.Client = *client
	if err := m.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	collection = m.Client.Database("engine").Collection("user_context")
}
