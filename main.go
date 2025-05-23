package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joaofilippe/pegtech/internal/application"
	"github.com/joaofilippe/pegtech/internal/application/api"
	locker_repositories "github.com/joaofilippe/pegtech/internal/application/repositories/locker"
	user_repositories "github.com/joaofilippe/pegtech/internal/application/repositories/user"
	"github.com/joaofilippe/pegtech/internal/application/services"
	"github.com/joaofilippe/pegtech/internal/infra/http"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
	"github.com/joho/godotenv"
)

func init() {
	// Carrega o arquivo .env se ele existir
	if err := godotenv.Load(); err != nil {
		log.Printf("Arquivo .env não encontrado: %v", err)
	}
}

func main() {

	mqttClient, err := mqtt.NewClient(
		os.Getenv("MQTT_BROKER"),
		os.Getenv("MQTT_PORT"),
		os.Getenv("MQTT_PROTOCOL"),
		os.Getenv("MQTT_USERNAME"),
		os.Getenv("MQTT_PASSWORD"),
		os.Getenv("MQTT_CLIENT_ID"),
		os.Getenv("MQTT_CA_CERT"),
	)
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %v", err)
	}

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
	userRepo := user_repositories.NewUserRepository(db)
	lockerRepo := locker_repositories.NewLockerRepository(db, mqttClient)

	// Initialize MQTT subscriber
	subscriber := mqtt.NewSubscriber(mqttClient)
	if err := subscriber.Start(); err != nil {
		log.Printf("Error starting MQTT subscriber: %v", err)
	}

	// Initialize services
	lockerService := services.NewLockerService(lockerRepo)
	userService := services.NewUserService(userRepo)

	application := application.NewApplication(lockerService, userService)

	// Create servers
	httpServer := http.NewHTTPServer()

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
