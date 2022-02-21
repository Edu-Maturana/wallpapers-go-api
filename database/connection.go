package database

import (
	"context"
	"log"
	"time"
	"wallpapers/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI = utils.GetEnv("MONGODB_ATLAS_URI")

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	} else {
		log.Println("Connected to database")
	}

	return client
}
