package routes

import (
	"net/http"

	"github.com/joaofilippe/pegtech/domain/iservices"
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
	e.GET("/packages/:trackingCode/pickup", r.getPackagePickupInfo)
	e.POST("/lockers/:lockerID/open", r.openLocker)
}

// registerPackage handles package registration
func (r *PackageRoutes) registerPackage(c echo.Context) error {
	var input struct {
		TrackingCode string `json:"trackingCode"`
		Size         string `json:"size"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	pkg, err := r.lockerService.RegisterPackage(input.TrackingCode, input.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, pkg)
}

// getPackagePickupInfo handles retrieval of package pickup information
func (r *PackageRoutes) getPackagePickupInfo(c echo.Context) error {
	trackingCode := c.Param("trackingCode")

	pickupInfo, err := r.lockerService.GetPackagePickupInfo(trackingCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, pickupInfo)
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

	if err := r.lockerService.OpenLocker(lockerID, input.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
