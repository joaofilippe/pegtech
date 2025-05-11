package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joaofilippe/pegtech/application"
	"github.com/joaofilippe/pegtech/application/repositories"
	"github.com/joaofilippe/pegtech/application/services"
	"github.com/joaofilippe/pegtech/infra/http"
	"github.com/joaofilippe/pegtech/infra/mqtt"
	"github.com/joaofilippe/pegtech/infra/repositories/memory"
)

func main() {
	// Initialize repositories
	lockerRepo := memory.NewLockerRepository()
	packageRepo := memory.NewPackageRepository()
	userRepo := repositories.NewUserRepository()

	// Initialize services
	lockerService := services.NewLockerService(lockerRepo, packageRepo)
	userService := services.NewUserService(userRepo)

	application := application.NewApplication(lockerService, userService)

	// Create servers
	httpServer := http.NewHTTPServer(application)
	mqttServer := mqtt.NewMQTTServer(lockerService)

	// Start MQTT server
	go func() {
		if err := mqttServer.Start(); err != nil {
			log.Fatalf("Error starting MQTT server: %v", err)
		}
	}()

	// Start HTTP server
	go func() {
		if err := httpServer.Start(":8080"); err != nil {
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
	if err := mqttServer.Shutdown(); err != nil {
		log.Printf("Error shutting down MQTT server: %v", err)
	}

	<-ctx.Done()
	log.Println("Servers shutdown complete")
}
