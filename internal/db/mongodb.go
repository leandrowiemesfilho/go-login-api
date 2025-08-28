package db

import (
	"context"
	"log"
	"time"

	"github.com/leandrowiemesfilho/login-api/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var mongoDB *MongoDB

func NewMongoDB(conf *configs.Config) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: conf.MongoUsername,
		Password: conf.MongoPassword,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoURI).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB is not reachable", err)
	}

	log.Println("Connected to MongoDB!")

	mongoDB = &MongoDB{
		client,
		client.Database(conf.MongoDB),
	}

	return mongoDB
}

func (m *MongoDB) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := m.Client.Disconnect(ctx)
	if err != nil {
		return
	}
}

func GetCollection(name string) *mongo.Collection {
	return mongoDB.Database.Collection(name)
}
