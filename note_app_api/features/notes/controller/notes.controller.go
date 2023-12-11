package controller

import (
	"net/http"
	"note_app_api/features/notes/service"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	data, err := service.Create(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new note.",
		"data":       data,
	})
}


