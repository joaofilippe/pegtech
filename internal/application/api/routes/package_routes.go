package routes

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	e.POST("/lockers/:lockerID/open", r.openLocker)
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

	pkg, err := r.lockerService.RegisterPackage(uuid.MustParse(input.UserID), input.LockerID, input.ExpiresAt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, pkg)
}

// openLocker handles locker opening for package pickup
func (r *PackageRoutes) openLocker(c echo.Context) error {
	lockerID := c.Param("lockerID")

	var input struct {
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	lockerIDInt, err := strconv.Atoi(lockerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid locker ID")
	}

	if err := r.lockerService.OpenLocker(lockerIDInt, input.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
