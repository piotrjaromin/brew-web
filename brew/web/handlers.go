package web

import (
	"fmt"
	"log"
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

func InitTemp(e *echo.Echo, t keg.Temperatures) {
	get := func(c echo.Context) error {
		return c.JSON(http.StatusOK, t.Get())
	}

	e.GET("/temperatures", get)
}

func InitTempControl(e *echo.Echo, tempControl keg.TempControl) {
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

func InitRecipes(e *echo.Echo, cook keg.Cook) {
	get := func(c echo.Context) error {
		return c.JSON(http.StatusOK, keg.RecipeStruct{})
	}

	post := func(c echo.Context) error {
		recipe := new(keg.RecipeStruct)
		if err := c.Bind(recipe); err != nil {
			return err
		}

		cook.Execute(recipe)
		return c.JSON(http.StatusOK, recipe)
	}

	delete := func(c echo.Context) error {
		cook.Stop()
		return c.NoContent(http.StatusNoContent)
	}

	e.GET("/recipes", get)
	e.POST("/recipes", post)
	e.DELETE("/recipes", delete)
}

func HandlerError(rw http.ResponseWriter, err error) {
	log.Fatal("Error while handling request. ", err.Error())
	rw.Write([]byte("error: " + err.Error()))
	rw.WriteHeader(http.StatusInternalServerError)
}
