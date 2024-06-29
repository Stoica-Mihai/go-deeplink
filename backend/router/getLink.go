package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetLinkInfoHandler(c echo.Context) error {
	fmt.Println("Getting link info")

	return nil
}
