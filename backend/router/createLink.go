package router

import (
	"fmt"
	"net/http"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating link")
}
