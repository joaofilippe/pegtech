package routes

import (
	"github.com/joaofilippe/pegtech/domain/iservices"
	"github.com/labstack/echo/v4"
)

// Router handles all API routes
type Router struct {
	userRoutes    *UserRoutes
	lockerRoutes  *LockerRoutes
	packageRoutes *PackageRoutes
}

// NewRouter creates a new instance of Router
func NewRouter(
	userService iservices.UserService,
	lockerService iservices.LockerService,
) *Router {
	return &Router{
		userRoutes:    NewUserRoutes(userService),
		lockerRoutes:  NewLockerRoutes(lockerService),
		packageRoutes: NewPackageRoutes(lockerService),
	}
}

// Setup configures all routes
func (r *Router) Setup(e *echo.Echo)  {
	// Register all routes
	r.userRoutes.Register(e)
	r.lockerRoutes.Register(e)
	r.packageRoutes.Register(e)
}
