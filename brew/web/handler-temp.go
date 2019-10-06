package web

import (
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/temperature"

	"github.com/labstack/echo/v4"
)

func InitTemp(e *echo.Echo, t temperature.Temperatures) {
	get := func(c echo.Context) error {
		point, err := t.Get()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, point)
	}

	e.GET("/temperatures", get)
}
