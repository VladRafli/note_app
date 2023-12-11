package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}