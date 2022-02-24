package services

import (
	"context"
	"fmt"
	"time"
	"wallpapers/database"
	"wallpapers/models"

	"github.com/rs/xid"
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
	wallpaper.ID = xid.New().String()
	wallpaper.Date = time.Now().Format("2006-01-02 15:04:05")
	_, err := collection.InsertOne(context.TODO(), wallpaper)
	if err != nil {
		return err
	}

	return nil
}

func DeleteWallpaper(id string) error {
	wallpaper, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return err
	}

	if wallpaper.DeletedCount == 0 {
		return fmt.Errorf("wallpaper not found")
	}

	return nil
}
