package routes

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	"github.com/labstack/echo/v4"
)

// LockerRoutes handles all locker-related routes
type LockerRoutes struct {
	lockerService iservices.LockerService
}

// NewLockerRoutes creates a new instance of LockerRoutes
func NewLockerRoutes(lockerService iservices.LockerService) *LockerRoutes {
	return &LockerRoutes{
		lockerService: lockerService,
	}
}

// Register registers all locker routes
func (r *LockerRoutes) Register(e *echo.Echo) {
	e.POST("/lockers", r.registerLocker)
	e.GET("/lockers/:id", r.getLocker)
	e.GET("/lockers/availables", r.getAvailableLockers)
	e.PUT("/lockers/:id/status", r.updateLockerStatus)
	e.GET("/lockers", r.listLockers)
	e.POST("/lockers/package", r.registerPackage)
	e.POST("/lockers/package/pickup", r.pickupPackage)
}

// registerLocker handles locker registration
func (r *LockerRoutes) registerLocker(c echo.Context) error {
	var input struct {
		ID    int   `json:"id"`
		Ports []int `json:"ports"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := r.lockerService.RegisterLocker(input.ID, input.Ports); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

// getAvailableLocker handles retrieval of available lockers by size
func (r *LockerRoutes) getAvailableLockers(c echo.Context) error {
	locker, err := r.lockerService.GetAvailableLockers()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, locker)
}

// getLocker handles locker retrieval by ID
func (r *LockerRoutes) getLocker(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid locker ID")
	}

	locker, err := r.lockerService.GetLocker(idInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, locker)
}

// updateLockerStatus handles locker status updates
func (r *LockerRoutes) updateLockerStatus(c echo.Context) error {
	id := c.Param("id")

	var input struct {
		Status string `json:"status"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid locker ID")
	}

	if err := r.lockerService.UpdateLockerStatus(idInt, entities.LockerStatus(input.Status)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// listLockers handles retrieval of all lockers
func (r *LockerRoutes) listLockers(c echo.Context) error {
	lockers, err := r.lockerService.ListLockers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, lockers)
}

// registerPackage handles package registration
func (r *LockerRoutes) registerPackage(c echo.Context) error {
	var input struct {
		UserID         string `json:"user_id"`
		ExpirationTime int    `json:"expiration_time"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID format")
	}

	packageCode, err := r.lockerService.RegisterPackage(userID, input.ExpirationTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"package_code": packageCode,
		"message":      "Package registered successfully",
	}

	return c.JSON(http.StatusCreated, response)
}

// pickupPackage handles package pickup
func (r *LockerRoutes) pickupPackage(c echo.Context) error {
	var input struct {
		PackageCode string `json:"package_code"`
		Password    string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := r.lockerService.PickupPackage(input.PackageCode, input.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
