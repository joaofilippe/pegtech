package routes

import (
	"net/http"

	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	"github.com/labstack/echo/v4"
)

// PackageRoutes handles all package-related routes
type PackageRoutes struct {
	lockerService iservices.LockerService
}

// NewPackageRoutes creates a new instance of PackageRoutes
func NewPackageRoutes(lockerService iservices.LockerService) *PackageRoutes {
	return &PackageRoutes{
		lockerService: lockerService,
	}
}

// Register registers all package routes
func (r *PackageRoutes) Register(e *echo.Echo) {
	e.POST("/packages", r.registerPackage)
}

// registerPackage handles package registration
func (r *PackageRoutes) registerPackage(c echo.Context) error {
	var input struct {
		UserID       string `json:"user_id"`
		ExpiresAt    int    `json:"expires_at"`
		LockerID     int    `json:"locker_id"`
		TrackingCode string `json:"tracking_code"`
		Size         string `json:"size"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	return c.JSON(http.StatusCreated, "ok")
}