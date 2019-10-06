package web

import (
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/recepies"

	"github.com/labstack/echo/v4"
)

func InitRecipes(e *echo.Echo, cook recepies.Cook) {
	get := func(c echo.Context) error {
		return c.JSON(http.StatusOK, recepies.RecipeStruct{})
	}

	post := func(c echo.Context) error {
		recipe := new(recepies.RecipeStruct)
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
