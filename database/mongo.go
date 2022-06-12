package database

import (
	"context"
	"fmt"
	"log"
	"loquigo/engine/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var databaseName string

var collection *mongo.Collection
var ctx = context.TODO()

func GetMongoConnection() *mongo.Database {
	return mongoClient.Database(databaseName)
}

func NewMongoDb(cfg infrastructure.DatabaseConfig, logger *infrastructure.Logger) MongoDB {
	databaseName = cfg.DbName()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
		cfg.DbUser(), cfg.DbPassword(), cfg.DbHost(), cfg.DbPort())
	return MongoDB{uri: uri, dbname: cfg.DbName(), logger: logger}
}

type MongoDB struct {
	uri    string
	dbname string
	logger *infrastructure.Logger
}

func (m *MongoDB) Connect() {
	m.logger.Info("Connecting to mongodb")
	clientOptions := options.Client().ApplyURI(m.uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	mongoClient = client
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	m.logger.Info("Connected to MongoDb")
}

func (m *MongoDB) Disconnect() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (m *MongoDB) Collection(name string) mongo.Collection {
	return *mongoClient.Database(m.dbname).Collection(name)
}
