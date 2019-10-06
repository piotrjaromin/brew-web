package web

import (
	"fmt"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/keg"

	"github.com/labstack/echo/v4"
)

func InitHeater(e *echo.Echo, kegControl keg.KegControl) {
	get := func(c echo.Context) error {
		power := kegControl.GetHeaterPower()
		return c.JSON(http.StatusOK, HeaterPower{power})
	}

	post := func(c echo.Context) error {
		power := new(HeaterPower)
		if err := c.Bind(power); err != nil {
			return err
		}

		kegControl.SetHeaterPower(power.Power)
		return c.JSON(http.StatusOK, power)
	}

	e.GET(fmt.Sprintf("/heaters"), get)
	e.POST(fmt.Sprintf("/heaters"), post)
}

type HeaterPower struct {
	Power float64 `json:"power"`
}
