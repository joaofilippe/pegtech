package routes

import (
	"net/http"
	"strconv"

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
}

// registerLocker handles locker registration
func (r *LockerRoutes) registerLocker(c echo.Context) error {
	var input struct {
		ID int `json:"id"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := r.lockerService.RegisterLocker(input.ID); err != nil {
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
