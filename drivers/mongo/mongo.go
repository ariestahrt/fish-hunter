package mongo_driver

import (
	"context"
	"fish-hunter/util"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mainDB *mongo.Database

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(util.GetConfig("MONGO_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, err
}

func SetClient(c *mongo.Client) {
	client = c
	mainDB = client.Database(util.GetConfig("MONGO_MAIN_DB"))
}

func GetDB() *mongo.Database {
	return mainDB
}

func GetCollection(name string) *mongo.Collection {
	return mainDB.Collection(name)
}