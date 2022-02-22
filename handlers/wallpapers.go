package handlers

import (
	"wallpapers/middlewares"
	"wallpapers/models"
	"wallpapers/services"

	"github.com/labstack/echo/v4"
)

func GetWallpapers(c echo.Context) error {
	wallpapers, err := services.GetWallpapers()
	if err != nil {
		return err
	}

	return c.JSON(200, wallpapers)
}

func CreateWallpaper(c echo.Context) error {
	var wallpaper models.Wallpaper

	adminError := middlewares.IsAdmin(c)
	if adminError != nil {
		return adminError
	}

	err := c.Bind(&wallpaper)
	if err != nil {
		return err
	}

	err = services.CreateWallpaper(wallpaper)
	if err != nil {
		return err
	}

	return c.JSON(200, "Wallpaper created successfully!")
}
