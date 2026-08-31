package mid

import (
	"context"
	"fmt"
	"time"

	"github.com/Joseeptessele/ardanlabs-service/foundation/logger"
	"github.com/Joseeptessele/ardanlabs-service/foundation/web"
)

func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	v := web.GetValues(ctx)

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	log.Info(ctx, "request started", "method", method, "path", path, "remoteAddr", remoteAddr)

	err := handler(ctx)

	log.Info(ctx, "request completed", "method", method, "path", path, "remoteAddr", remoteAddr,
		"status", v.StatusCode, "since", time.Since(v.Now))

	return err
}
