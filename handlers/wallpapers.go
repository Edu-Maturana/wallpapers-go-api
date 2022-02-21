package handlers

import (
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
