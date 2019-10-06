package web

import (
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/temperature"

	"github.com/labstack/echo/v4"
)

func InitTempControl(e *echo.Echo, tempControl temperature.TempControl) {
	get := func(c echo.Context) error {
		return c.JSON(http.StatusOK, Temp{tempControl.GetTemp()})
	}

	delete := func(c echo.Context) error {
		tempControl.Stop()
		return c.NoContent(http.StatusNoContent)
	}

	post := func(c echo.Context) error {
		temp := new(Temp)

		if err := c.Bind(temp); err != nil {
			return err
		}

		tempControl.KeepTemp(temp.Value)
		return c.JSON(http.StatusOK, temp)
	}

	e.GET("/temperatures/control", get)
	e.POST("/temperatures/control", post)
	e.DELETE("/temperatures/control", delete)
}
