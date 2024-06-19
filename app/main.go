package main

import "net/http"

func main() {
	router := http.NewServeMux()

	http.ListenAndServe(":9000", router)
}
