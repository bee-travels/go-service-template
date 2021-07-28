package main

import (
	"github.com/{PROJECT_PATH}/{{.ServiceNamePill}}/wrappers/server" // Update
	instana "github.com/instana/go-sensor"
)

func main() {
	if err := server.Start("{{.ServiceNamePill}}", initializeRouter); err != nil {
		panic(err)
	}
}

func initializeRouter(router server.PathRouter, _ *instana.Sensor) {
	router.Path("/api/v1/{{.Route}}", func(router server.PathRouter) {
		// path: /api/v1/{{.Route}}
		router.Get(getData())
	})
}
