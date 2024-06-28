package router

import (
	"net/http"
)

var FileServer = http.FileServer(http.Dir(STATICPATH))
