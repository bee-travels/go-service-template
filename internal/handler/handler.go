package handler

import (
	"net/http"

	"{{.ServiceNamePill}}/internal/data"
	"{{.ServiceNamePill}}/wrappers/server"
)

func GetData() server.RequestHandler {
	return func(ctx server.RequestContext) {
		err = nil
		if err != nil {
			server.Response(ctx, http.StatusForbidden, data.Error{
				Error: err.Error(),
			})
			return
		}
		server.Response(ctx, http.StatusOK, data.CustomType{
			Implement: "Implement me",
		})
	}
}
