package services

import (
	"context"
	"fmt"
	"wallpapers/database"
	"wallpapers/models"

	"go.mongodb.org/mongo-driver/bson"
)

var client = database.Connect()
var collection = client.Database("goproject").Collection("wallpapers")

func GetWallpapers() ([]models.Wallpaper, error) {
	var wallpapers []models.Wallpaper

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var wallpaper models.Wallpaper
		err := cursor.Decode(&wallpaper)
		if err != nil {
			return nil, err
		}

		wallpapers = append(wallpapers, wallpaper)
	}

	return wallpapers, nil
}

func CreateWallpaper(wallpaper models.Wallpaper) error {
	_, err := collection.InsertOne(context.TODO(), wallpaper)
	if err != nil {
		return err
	}

	return nil
}
