package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	URI string
}

func NewMongoConfig() *MongoDBConfig {
	uri := os.Getenv("MONGODB_URL")

	if uri == "" {
		panic("You must set your 'MONGODB_URL' environmental variable. Check the .env.example file")
	}
	return &MongoDBConfig{
		URI: uri,
	}
}

func (c *MongoDBConfig) Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.URI))
	defer c.Disconnect(client)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func (c *MongoDBConfig) Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
