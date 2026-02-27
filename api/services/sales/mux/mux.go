// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/makxtr/scratch-service/api/services/api/mid"
	"github.com/makxtr/scratch-service/api/services/sales/route/sys/checkapi"

	"github.com/makxtr/scratch-service/foundation/logger"
	"github.com/makxtr/scratch-service/foundation/web"
)

func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics())

	checkapi.Routes(mux)

	return mux
}
