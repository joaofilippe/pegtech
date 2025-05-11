package api

import (
	"github.com/joaofilippe/pegtech/application"
	"github.com/joaofilippe/pegtech/application/api/routes"
	"github.com/joaofilippe/pegtech/infra/http"
	"github.com/labstack/echo/v4"
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
	return &Api{
		application: application,
		httpServer:  http.NewHTTPServer(application),
	}
}

func (a *Api) RegisterRoutes(e *echo.Echo) {
	a.router.Setup(e)
}