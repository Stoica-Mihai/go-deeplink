package utils

import (
	"log"
	"math/rand"
	"net/http"
)

const (
	CHARSET      = "abcdefghijklmnopqrstuvwxyz-_+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CHARSETLEN   = len(CHARSET)
	DEEPLINKSIZE = 15
)

func GenerateDeepLink() string {

	var result string

	for DEEPLINKSIZE != len(result) {
		index := rand.Intn(CHARSETLEN)
		result += string(CHARSET[index])
	}

	return result
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
