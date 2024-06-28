package router

import (
	"log"
	"net/http"
)

func GetRouter() *http.ServeMux {
	router := http.NewServeMux()
	// routerLogging := utils.LoggingMiddleware(router)

	router.Handle("/", FileServer)

	router.HandleFunc("POST /create", CreateLink)

	router.HandleFunc("GET /link/{id}", GetLinkInfo)

	router.HandleFunc("DELETE /link/{id}", DeleteLink)

	return router
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
