package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type VersionResponse struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

func InitVersion(e *echo.Echo, version VersionResponse) {
	get := func(c echo.Context) error {
		return c.JSON(http.StatusOK, version)
	}

	e.GET("/version", get)
}
