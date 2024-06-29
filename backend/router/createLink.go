package router

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/publicsuffix"
)

const (
	CHARSET         = "abcdefghijklmnopqrstuvwxyz-_+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CHARSETLEN      = len(CHARSET)
	DEEPLINKSIZE    = 15
	JSONCONTENTTYPE = "application/json"
)

type CreateLinkResponse struct {
	Deeplink string `json:"deeplink"`
	Redirect string `json:"redirect"`
}

type CreateLinkError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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

	createLinkResponse.Deeplink = generateDeepLink()

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

func generateDeepLink() string {

	var result string

	for DEEPLINKSIZE != len(result) {
		index := rand.Intn(CHARSETLEN)
		result += string(CHARSET[index])
	}

	return result
}
