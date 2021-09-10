package main

import (
	"{{.ServiceNamePill}}/internal/handler"
	"{{.ServiceNamePill}}/wrappers/server"

	instana "github.com/instana/go-sensor"
)

func main() {
	if err := server.Start("{{.ServiceNamePill}}", initializeRouter); err != nil {
		panic(err)
	}
}

func initializeRouter(router server.PathRouter, _ *instana.Sensor) {
	router.Path("/api/v1/{{.Route}}", func(router server.PathRouter) {
		/**
		* GET /api/v1/{{.Route}}
		* @tag {{.Route}}
		* @summary Get data
		* @description {{.ServiceNameTitle}} data
		* @response 200 - OK
		* @response 500 - Internal Server Error
		*/
		router.Get(handler.GetData())
	})
}
