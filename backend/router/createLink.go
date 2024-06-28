package router

import (
	"go-deeplink/backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Deeplink utils.Deeplink `json:"deeplink"`
	Redirect string         `json:"redirect"`
}

func CreateLinkHandler(c echo.Context) error {
	CheckHeaders(c)

	response := new(Response)

	if err := c.Bind(response); err != nil {
		return c.String(http.StatusBadRequest, "400 - Malformed request body")
	}

	response.Deeplink = utils.GenerateDeepLink()

	return c.JSON(http.StatusOK, response)
}

func CheckHeaders(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.String(http.StatusBadRequest, "400 - Missing or invalid Content-Type header\n")
	}
	return nil
}
