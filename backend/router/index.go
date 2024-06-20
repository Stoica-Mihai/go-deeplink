package router

import (
	"fmt"
	"go-deeplink/backend/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(utils.GenerateDeepLink())
}
