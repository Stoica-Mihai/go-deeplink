package router

import (
	"fmt"
	"go-deeplink/backend/utils"
	"net/http"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating link")
	fmt.Println(utils.GenerateDeepLink())
}
