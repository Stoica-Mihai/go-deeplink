package router

import (
	"go-deeplink/backend/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/publicsuffix"
)

type CreateLinkResponse struct {
	Deeplink utils.Deeplink `json:"deeplink"`
	Redirect string         `json:"redirect"`
}

type CreateLinkError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const JSONCONTENTTYPE = "application/json"

func CreateLinkHandler(c echo.Context) error {

	if c.Request().Header.Get("Content-Type") != JSONCONTENTTYPE {
		return c.JSON(http.StatusBadRequest, &CreateLinkError{
			Status:  http.StatusBadRequest,
			Message: "Missing or invalid Content-Type header",
		})
	}

	createLinkResponse := &CreateLinkResponse{}

	if err := c.Bind(createLinkResponse); err != nil {
		return c.JSON(http.StatusBadRequest, &CreateLinkError{
			Status:  http.StatusBadRequest,
			Message: "Malformed request body",
		})
	}

	if !isValidURL(createLinkResponse.Redirect) {
		return c.JSON(http.StatusBadRequest, &CreateLinkError{
			Status:  http.StatusBadRequest,
			Message: "Url is not valid",
		})
	}

	createLinkResponse.Deeplink = utils.GenerateDeepLink()

	return c.JSON(http.StatusOK, createLinkResponse)
}

func isValidURL(urlStr string) bool {
	if !strings.HasPrefix(urlStr, "http://") ||
		!strings.HasPrefix(urlStr, "https://") ||
		!strings.HasPrefix(urlStr, "www.") {
		urlStr = "http://" + urlStr
	}

	_, err := url.Parse(urlStr)

	if err != nil {
		return false
	}

	if _, ok := publicsuffix.PublicSuffix(urlStr); !ok {
		return false
	}

	return true
}
