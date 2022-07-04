package domain

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Db *mongo.Database
)

func ConnMongoDB() {
	uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}
	Db = client.Database("University")
	fmt.Println("Successfuly Connected to the mongodb")
}
