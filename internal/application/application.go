package application

import (
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
)

type Application struct {
	UserService   iservices.UserService
	LockerService iservices.LockerService
}

func NewApplication(
	lockerService iservices.LockerService,
	userService iservices.UserService,
) *Application {
	app := &Application{
		LockerService: lockerService,
		UserService:   userService,
	}

	app.init()
	return app
}

func (a *Application) init() {
		// Register 10 initial lockers
		for i := 1; i <= 10; i++ {
			err := a.LockerService.RegisterLocker(i)
			if err != nil {
				// Log error but continue
				continue
			}
		}
}
