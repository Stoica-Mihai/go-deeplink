package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndexHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "main", nil)
}

func GetStatsHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "stats", nil)
}
