package routes

import (
	"net/http"

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
	e.GET("/lockers/available/:size", r.getAvailableLocker)
	e.GET("/lockers/:id", r.getLocker)
	e.PUT("/lockers/:id/status", r.updateLockerStatus)
	e.GET("/lockers", r.listLockers)
}

// registerLocker handles locker registration
func (r *LockerRoutes) registerLocker(c echo.Context) error {
	var input struct {
		ID   string `json:"id"`
		Size string `json:"size"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := r.lockerService.RegisterLocker(input.ID, input.Size); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

// getAvailableLocker handles retrieval of available lockers by size
func (r *LockerRoutes) getAvailableLocker(c echo.Context) error {
	size := c.Param("size")

	locker, err := r.lockerService.GetAvailableLocker(size)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, locker)
}

// getLocker handles locker retrieval by ID
func (r *LockerRoutes) getLocker(c echo.Context) error {
	id := c.Param("id")

	locker, err := r.lockerService.GetLocker(id)
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

	if err := r.lockerService.UpdateLockerStatus(id, entities.LockerStatus(input.Status)); err != nil {
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
