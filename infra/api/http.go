package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	lockerusecases "github.com/joaofilippe/pegtech/domain/usecases/locker"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	echo          *echo.Echo
	lockerUseCase *lockerusecases.LockerUseCase
}

func NewHTTPServer(lockerUseCase *lockerusecases.LockerUseCase) *HTTPServer {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &HTTPServer{
		echo:          e,
		lockerUseCase: lockerUseCase,
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

type PackageRegistrationRequest struct {
	TrackingCode string `json:"tracking_code"`
	Size         string `json:"size"`
}

func (s *HTTPServer) registerPackage(c echo.Context) error {
	var req PackageRegistrationRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	pkg, err := s.lockerUseCase.RegisterPackage(req.TrackingCode, req.Size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, pkg)
}

func (s *HTTPServer) getLockers(c echo.Context) error {
	// This would need to be implemented in the LockerUseCase
	// For now, returning a placeholder response
	return c.JSON(http.StatusOK, map[string]string{"message": "List of lockers"})
}

func (s *HTTPServer) getPackageInfo(c echo.Context) error {
	trackingCode := c.Param("trackingCode")
	if trackingCode == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tracking code is required"})
	}

	pickupInfo, err := s.lockerUseCase.GetPackagePickupInfo(trackingCode)
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
