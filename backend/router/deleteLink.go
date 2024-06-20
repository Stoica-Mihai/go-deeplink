package router

import (
	"fmt"
	"net/http"
)

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting link")
}
