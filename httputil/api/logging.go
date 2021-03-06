package api

import (
	"context"

	logging "github.com/puppetlabs/horsehead/v2/logging"
)

var (
	logger = logging.Builder().At("horsehead", "httputil", "api")
)

func log(ctx context.Context) logging.Logger {
	return logger.With(ctx).Build()
}
