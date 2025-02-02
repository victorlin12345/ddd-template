package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mClient *MongoClient

type MongoClient struct {
	Client *mongo.Client
}

func NewMongoClient() *MongoClient {
	if mClient == nil {
		// Connect to MongoDB
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			fmt.Println("failed to connect to mongo", err)
			log.Fatal(err)
		}

		mClient = &MongoClient{
			Client: client,
		}
	}
	return mClient
}

func (mc *MongoClient) Start(ctx context.Context) error {
	fmt.Println("mongo client start")
	return mc.Client.Ping(ctx, nil)
}

func (mc *MongoClient) Stop(ctx context.Context) error {
	fmt.Println("mongo client stop")
	return mc.Client.Disconnect(ctx)
}

func GetMongoClient() *MongoClient {
	return mClient
}
