package router

import (
	"net/http"
)

const FRONTENDPATH = "/app/frontend/"

var FileServer = http.FileServer(http.Dir(FRONTENDPATH))
