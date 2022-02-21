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

	error := middlewares.IsAdmin(c)
	if error != nil {
		return error
	}

	if err := c.Bind(&wallpaper); err != nil {
		return err
	}

	err := services.CreateWallpaper(wallpaper)
	if err != nil {
		return err
	}

	return c.JSON(200, wallpaper)
}
