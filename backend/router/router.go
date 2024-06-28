package router

import (
	"path"

	"github.com/labstack/echo/v4"
)

const FRONTENDPATH = "/app/frontend"

var STATICPATH = path.Join(FRONTENDPATH, "static")
var VIEWSPATH = path.Join(FRONTENDPATH, "views")

func GetRouter() *echo.Echo {
	router := echo.New()

	router.Static("/static", STATICPATH)

	router.POST("/create", CreateLinkHandler)

	router.GET("/link/:id", GetLinkInfo)

	router.DELETE("/link/:id", DeleteLink)

	return router
}
