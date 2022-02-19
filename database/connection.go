package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("to do: add uri"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	} else {
		log.Println("Connected to database")
	}
	defer client.Disconnect(ctx)
}
