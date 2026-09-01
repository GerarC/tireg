package rest

import "net/http"

func HttpRestAPIInitializer(handler http.Handler, port string) error {
	return http.ListenAndServe(":"+port, handler)
}
