package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	lockerusecases "github.com/joaofilippe/pegtech/domain/usecases/locker"
	"github.com/joaofilippe/pegtech/infra/api"
	"github.com/joaofilippe/pegtech/infra/mqtt"
	"github.com/joaofilippe/pegtech/infra/repositories/memory"
)

func main() {
	// Initialize repositories
	lockerRepo := memory.NewLockerRepository()
	packageRepo := memory.NewPackageRepository()

	// Initialize use case
	lockerUseCase := lockerusecases.NewLockerUseCase(lockerRepo, packageRepo)

	// Register some initial lockers
	lockerUseCase.RegisterLocker("L001", "small")
	lockerUseCase.RegisterLocker("L002", "medium")
	lockerUseCase.RegisterLocker("L003", "large")

	// Create servers
	httpServer := api.NewHTTPServer(lockerUseCase)
	mqttServer := mqtt.NewMQTTServer(lockerUseCase)

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
