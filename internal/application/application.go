package application

import (
	"log"

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

	err := app.startSubscriptions()
	if err != nil {
		log.Fatalf("Error starting subscriptions: %v", err)
	}

	// err = app.init()
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	return app
}

func (a *Application) init() error {
	for i := 1; i <= 4; i++ {
		err := a.LockerService.RegisterLocker(i, []int{1, 2, 3, 4})
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) startSubscriptions() error {
	_, err := a.LockerService.StartPackagePickupSubscription()
	if err != nil {
		return err
	}

	return a.LockerService.StartRegisterPackageSubscription()
}
