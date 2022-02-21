package main

import (
	"wallpapers/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/wallpapers", handlers.GetWallpapers)

	app.Start(":8080")
}
