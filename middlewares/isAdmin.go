package middlewares

import (
	"wallpapers/utils"

	"github.com/labstack/echo/v4"
)

var adminEmail = utils.GetEnv("ADMIN_EMAIL")

func IsAdmin(ctx echo.Context) error {
	email := ctx.FormValue("email")

	if email != adminEmail {
		return echo.ErrForbidden
	}

	return nil
}
