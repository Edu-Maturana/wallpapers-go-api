package main

import (
	"net/http"
	"wallpapers/database"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	database.Connect()
	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})
	app.Start(":8080")
}
