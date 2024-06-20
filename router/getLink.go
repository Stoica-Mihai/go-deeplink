package router

import (
	"fmt"
	"net/http"
)

func GetLinkInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting link info")
}
