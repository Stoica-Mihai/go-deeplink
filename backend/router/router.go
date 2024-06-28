package router

import (
	"path"

	"github.com/labstack/echo/v4"
)

const FRONTENDPATH = "/app/frontend"

var STATICPATH = path.Join(FRONTENDPATH, "static")
var VIEWSPATH = path.Join(FRONTENDPATH, "views")

// func GetRouter() *http.ServeMux {
// 	router := http.NewServeMux()

// 	// router.HandleFunc("/", GetView("main.html"))
// 	// router.HandleFunc("/stats", GetView("stats.html"))

// 	router.Handle("/static", FileServer)

// 	router.HandleFunc("POST /create", CreateLink)

// 	router.HandleFunc("GET /link/{id}", GetLinkInfo)

// 	router.HandleFunc("DELETE /link/{id}", DeleteLink)

// 	return router
// }

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("[%s] %s\n", r.Method, r.URL.String())
// 		next.ServeHTTP(w, r)
// 	})
// }

func GetRouter() *echo.Echo {
	router := echo.New()

	router.Static("/static", STATICPATH)

	router.POST("/create", CreateLinkHandler)

	router.GET("/link/:id", GetLinkInfo)

	router.DELETE("/link/:id", DeleteLink)

	return router
}
