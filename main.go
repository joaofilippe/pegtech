package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joaofilippe/pegtech/application"
	"github.com/joaofilippe/pegtech/application/api"
	"github.com/joaofilippe/pegtech/application/repositories"
	"github.com/joaofilippe/pegtech/application/services"
	"github.com/joaofilippe/pegtech/infra/http"
	"github.com/joaofilippe/pegtech/infra/repositories/database"
	"github.com/joaofilippe/pegtech/infra/repositories/memory"
)

func main() {

	db, err := database.NewPostgresDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize repositories
	lockerRepo := memory.NewLockerRepository()
	packageRepo := memory.NewPackageRepository()
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	lockerService := services.NewLockerService(lockerRepo, packageRepo)
	userService := services.NewUserService(userRepo)

	application := application.NewApplication(lockerService, userService)

	// Create servers
	httpServer := http.NewHTTPServer(application)

	api := api.NewApi(application, httpServer)
		// Start MQTT server


	// Start HTTP server
	go func() {
		if err := api.Start(); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the servers
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown servers with context
	if err := httpServer.Shutdown(); err != nil {
		log.Printf("Error shutting down HTTP server: %v", err)
	}


	<-ctx.Done()
	log.Println("Servers shutdown complete")
}
