package main

import (
	"wallpapers/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	app := echo.New()

	// CORS
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	app.GET("/wallpapers", handlers.GetWallpapers)
	app.POST("/wallpapers", handlers.CreateWallpaper)
	app.DELETE("/wallpapers/:id", handlers.DeleteWallpaper)

	// Start server
	app.Start(":8080")
}
