package application

import "github.com/joaofilippe/pegtech/domain/iservices"

type Application struct {
	userService   iservices.UserService
	lockerService iservices.LockerService
}

func NewApplication(
	lockerService iservices.LockerService,
	userService iservices.UserService,
) *Application {
	return &Application{
		lockerService: lockerService,
		userService:   userService,
	}
}
