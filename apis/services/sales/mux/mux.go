package mux

import (
	"os"

	"github.com/Joseeptessele/ardanlabs-service/apis/services/sales/route/sys/checkapi"
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
