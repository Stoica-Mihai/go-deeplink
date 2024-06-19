package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessing index")
	})

	router.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("creating dynamic link")
	})

	router.HandleFunc("GET /link/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("creating dynamic link")
	})

	router.HandleFunc("DELETE /link/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("deleting dynamic link")
	})

	http.ListenAndServe(":9000", router)
}
