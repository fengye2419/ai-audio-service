package common

import (
	"net/http"
)

func Middlewares() []func(http.Handler) http.Handler {
	var handlers = []func(http.Handler) http.Handler{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				next.ServeHTTP(resp, req)
			})
		},
	}
	handlers = append(handlers, LoggerHandler())

	return handlers
}
