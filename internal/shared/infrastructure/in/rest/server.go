package rest

import "net/http"

func HttpRestAPIInitializer(mux *http.ServeMux, port string) error {
	return http.ListenAndServe(":"+port, mux)
}
