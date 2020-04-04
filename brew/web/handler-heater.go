package web

import (
	"fmt"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/keg"

	"github.com/labstack/echo/v4"
)

type BadRequest struct {
	Message string `json:"message"`
}

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

		powerVal := power.Power
		if powerVal < 0 || powerVal > 1 {
			return c.JSON(http.StatusBadRequest, BadRequest{
				Message: fmt.Sprintf("Power must be between 0 and 1, got %f", powerVal),
			})
		}

		kegControl.SetHeaterPower(powerVal)
		return c.JSON(http.StatusOK, power)
	}

	e.GET("/heaters", get)
	e.POST("/heaters", post)
}

type HeaterPower struct {
	Power float64 `json:"power"`
}
