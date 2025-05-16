package api

import (
	"github.com/joaofilippe/pegtech/internal/application"
	"github.com/joaofilippe/pegtech/internal/application/api/routes"
	"github.com/joaofilippe/pegtech/internal/infra/http"
)

type Api struct {
	application *application.Application
	router      *routes.Router
	httpServer  *http.HTTPServer
}

func NewApi(
	application *application.Application,
	httpServer *http.HTTPServer,
) *Api {
	a := &Api{
		application: application,
		httpServer:  http.NewHTTPServer(),
	}
	a.registerRoutes()
	return a
}

func (a *Api) Start() error {
	return a.httpServer.Start(":8080")
}

func (a *Api) registerRoutes() {
	a.router = routes.NewRouter(
		a.application.UserService,
		a.application.LockerService,
	)

	a.router.Setup(a.httpServer)
}
