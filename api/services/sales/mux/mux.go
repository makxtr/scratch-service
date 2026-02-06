// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/makxtr/scratch-service/api/services/sales/route/sys/checkapi"
	"github.com/makxtr/scratch-service/foundation/web"
)

func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
