package mux

import (
	"net/http"

	"github.com/Joseeptessele/ardanlabs-service/apis/services/sales/route/sys/checkapi"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
