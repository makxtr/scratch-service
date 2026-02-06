// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"

	"github.com/makxtr/scratch-service/api/services/sales/route/sys/checkapi"
)

func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
