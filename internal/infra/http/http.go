package http

import (
	"context"
	"time"

	"github.com/joaofilippe/pegtech/application"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	echo *echo.Echo
}

func NewHTTPServer(application *application.Application) *HTTPServer {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &HTTPServer{e}

	return server
}

func (s *HTTPServer) Echo() *echo.Echo {
	return s.echo
}

func (s *HTTPServer) Start(address string) error {
	return s.echo.Start(address)
}

func (s *HTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.echo.Shutdown(ctx)
}
