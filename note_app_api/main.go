package main

import (
	"note_app_api/helper"
	"note_app_api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	routes.Init(app)
	helper.LoadEnvironment(app)
}