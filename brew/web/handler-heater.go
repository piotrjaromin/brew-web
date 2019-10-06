package web

import (
	"fmt"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/keg"

	"github.com/labstack/echo/v4"
)

func InitHeater(e *echo.Echo, heater keg.Heater, kegControl keg.KegControl) {
	get := func(c echo.Context) error {
		state := kegControl.HeaterState(heater)
		return c.JSON(http.StatusOK, HeaterState{state})
	}

	post := func(c echo.Context) error {
		state := new(HeaterState)
		if err := c.Bind(state); err != nil {
			return err
		}

		kegControl.SetHeaterState(heater, state.State)
		return c.JSON(http.StatusOK, state)
	}

	e.GET(fmt.Sprintf("/heaters/%d", heater), get)
	e.POST(fmt.Sprintf("/heaters/%d", heater), post)
}
