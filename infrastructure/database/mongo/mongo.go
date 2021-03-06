package mongo

import (
	"context"
	"fmt"
	"log"
	"loquigo/engine/pkg/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var ctx = context.TODO()

func NewMongoDb(cfg infrastructure.DatabaseConfig, logger *infrastructure.Logger) MongoDB {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
		cfg.DbUser(), cfg.DbPassword(), cfg.DbHost(), cfg.DbPort())
	return MongoDB{uri: uri, dbname: cfg.DbName(), logger: logger}
}

type MongoDB struct {
	client mongo.Client
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
	m.client = *client
	if err := m.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	m.logger.Info("Connected to MongoDb")
}

func (m *MongoDB) Disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (m *MongoDB) Collection(name string) mongo.Collection {
	return *m.client.Database(m.dbname).Collection(name)
}
