package router

import (
	"go-deeplink/backend/utils"
	"net/http"
)

func GetRouter() http.Handler {
	router := http.NewServeMux()
	routerLogging := utils.LoggingMiddleware(router)

	router.Handle("/", FileServer)

	router.HandleFunc("POST /create", CreateLink)

	router.HandleFunc("GET /link/{id}", GetLinkInfo)

	router.HandleFunc("DELETE /link/{id}", DeleteLink)

	return routerLogging
}
