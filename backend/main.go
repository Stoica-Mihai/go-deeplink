package main

import (
	"fmt"
	"go-deeplink/router"
	"net/http"
)

const (
	IP   = "0.0.0.0"
	PORT = "9000"
)

func main() {
	ServerIPAndPort := IP + ":" + PORT
	fmt.Println("Listen on", ServerIPAndPort)

	http.ListenAndServe(ServerIPAndPort, router.GetRouter())
}
