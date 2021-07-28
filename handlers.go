package main

import (
	"github.com/{PROJECT_PATH}/{{.ServiceNamePill}}/wrappers/server" // Update
	"net/http"
)

func getData() server.RequestHandler {
	return func(ctx server.RequestContext) {
		err = nil
		if err != nil {
			server.Response(ctx, http.StatusForbidden, Error{
				Error: err.Error(),
			})
			return
		}
		server.Response(ctx, http.StatusOK, CustomType{
			Implement: "Implement me",
		})
	}
}
