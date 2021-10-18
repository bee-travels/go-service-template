package handler

import (
	"net/http"

	"{{.ServiceNamePill}}/internal/data"
	"{{.ServiceNamePill}}/wrappers/server"
)

func GetData() server.RequestHandler {
	return func(ctx server.RequestContext) {
		server.Response(ctx, http.StatusOK, data.CustomType{
			Implement: "Implement me",
		})
	}
}
