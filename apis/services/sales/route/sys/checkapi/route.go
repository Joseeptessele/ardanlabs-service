package checkapi

import (
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
}
