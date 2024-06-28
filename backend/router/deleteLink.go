package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func DeleteLink(c echo.Context) error {
	fmt.Println("Deleting link")

	return nil
}
