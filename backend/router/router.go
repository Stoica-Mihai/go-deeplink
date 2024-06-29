package router

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

const FRONTENDPATH = "/app/frontend/"

const STATICPATH = FRONTENDPATH + "static/"
const VIEWSPATH = FRONTENDPATH + "views/"

func GetRouter() *echo.Echo {
	router := echo.New()
	router.HideBanner = true

	t := NewTemplate(template.Must(template.ParseGlob(VIEWSPATH + "*.html")))

	router.Renderer = t

	router.GET("/", GetIndexHandler)
	router.GET("/stats", GetStatsHandler)

	router.Static("/static", STATICPATH)

	router.POST("/create", CreateLinkHandler)

	router.GET("/link/:id", GetLinkInfoHandler)

	router.DELETE("/link/:id", DeleteLinkHandler)

	return router
}
