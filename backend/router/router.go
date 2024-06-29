package router

import (
	"github.com/labstack/echo/v4"
)

const FRONTENDPATH = "/app/frontend/"

const STATICPATH = FRONTENDPATH + "static/"
const VIEWSPATH = FRONTENDPATH + "views/"

type (
	Router struct {
		*echo.Echo
	}

	RouterConfig struct {
		HideBanner bool
		Renderer   *Template
	}
)

func NewRouter(config *RouterConfig) *Router {
	router := &Router{
		Echo: echo.New(),
	}

	router.HideBanner = config.HideBanner
	router.Renderer = config.Renderer

	return router
}

func GetRouter() *Router {
	router := NewRouter(&RouterConfig{
		HideBanner: true,
		Renderer:   NewTemplate(VIEWSPATH + "*.html"),
	})

	router.Static("/static", STATICPATH)

	router.GET("/", GetIndexHandler)

	router.GET("/stats", GetStatsHandler)

	router.POST("/create", CreateLinkHandler)

	router.GET("/link/:id", GetLinkInfoHandler)

	router.DELETE("/link/:id", DeleteLinkHandler)

	return router
}
