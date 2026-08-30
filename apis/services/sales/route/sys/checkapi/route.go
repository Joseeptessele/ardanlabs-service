package checkapi

import (
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

func Routes(mux *web.App) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
}
