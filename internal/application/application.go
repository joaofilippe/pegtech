package application

import "github.com/joaofilippe/pegtech/internal/domain/iservices"

type Application struct {
	UserService   iservices.UserService
	LockerService iservices.LockerService
}

func NewApplication(
	lockerService iservices.LockerService,
	userService iservices.UserService,
) *Application {
	return &Application{
		LockerService: lockerService,
		UserService:   userService,
	}
}
