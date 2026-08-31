package checkapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/Joseeptessele/ardanlabs-service/app/api/errs"
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func testError(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "This message is trused")
	}

	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
