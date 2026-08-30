package mux

import (
	"os"

	"github.com/Joseeptessele/ardanlabs-service/apis/services/sales/route/sys/checkapi"
	"github.com/Joseeptessele/ardanlabs-service/app/api/mid"
	"github.com/Joseeptessele/ardanlabs-service/foundation/logger"
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log))

	checkapi.Routes(mux)

	return mux
}
