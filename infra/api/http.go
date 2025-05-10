package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/joaofilippe/pegtech/application/models"
	"github.com/joaofilippe/pegtech/application/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	echo          *echo.Echo
	lockerService *services.LockerService
}

func NewHTTPServer(lockerService *services.LockerService) *HTTPServer {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &HTTPServer{
		echo:          e,
		lockerService: lockerService,
	}

	// Employee routes
	employee := e.Group("/employee")
	employee.POST("/package", server.registerPackage)
	employee.GET("/lockers", server.getLockers)

	// Customer routes
	customer := e.Group("/customer")
	customer.GET("/package/:trackingCode", server.getPackageInfo)

	return server
}

func (s *HTTPServer) registerPackage(c echo.Context) error {
	var reg models.PackageRegistration
	if err := json.NewDecoder(c.Request().Body).Decode(&reg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	pkg, err := s.lockerService.RegisterPackage(reg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, pkg)
}

func (s *HTTPServer) getLockers(c echo.Context) error {
	// This would need to be implemented in the LockerService
	// For now, returning a placeholder response
	return c.JSON(http.StatusOK, map[string]string{"message": "List of lockers"})
}

func (s *HTTPServer) getPackageInfo(c echo.Context) error {
	trackingCode := c.Param("trackingCode")
	if trackingCode == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tracking code is required"})
	}

	pickupInfo, err := s.lockerService.GetPackagePickupInfo(trackingCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, pickupInfo)
}

func (s *HTTPServer) Start(address string) error {
	return s.echo.Start(address)
}

func (s *HTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.echo.Shutdown(ctx)
}
