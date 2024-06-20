package router

import "net/http"

func GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Index)

	router.HandleFunc("POST /create", CreateLink)

	router.HandleFunc("GET /link/{id}", GetLinkInfo)

	router.HandleFunc("DELETE /link/{id}", DeleteLink)

	return router
}
