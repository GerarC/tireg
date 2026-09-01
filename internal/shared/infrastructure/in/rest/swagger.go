package rest

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/util/constant"
)

func RegisterSwaggerRoutes(mux *http.ServeMux) {
	mux.Handle(constant.SwaggerRoutePrefix, httpSwagger.Handler())
}
