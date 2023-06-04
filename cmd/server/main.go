package main

import (
	"FACEIT-coding-test/config"
	publisher "FACEIT-coding-test/internal/infrastructure/publisher"
	"FACEIT-coding-test/internal/service"
	"log"
	httpNet "net/http"

	_ "github.com/lib/pq"

	"FACEIT-coding-test/internal/infrastructure/database/postgresql"
	"FACEIT-coding-test/internal/infrastructure/http"
	"FACEIT-coding-test/internal/infrastructure/logger"
)

func main() {
	// Set up logging
	defaultLogger := logger.DefaultLogger{}

	// Set up database connection
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// Create a new user repository and service
	userRepo, _ := postgresql.NewUserRepositoryPostgres(cfg.Database.ConnectionString)
	pb, err := publisher.NewRabbitMQPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Error configuring publisher: %v", err)
	}

	userService := service.NewUserService(userRepo, pb)

	userHandler := http.NewUserHandler(userService)

	// Create a new user router
	userRouter := http.NewRouter(userHandler)

	// Wrap the router with middleware
	loggingHandler := http.LoggingMiddleware(userRouter)
	recoveryHandler := http.RecoverMiddleware(loggingHandler)

	// Start the HTTP server
	addr := ":8080"
	defaultLogger.Infof("Starting server on %s", addr)
	if err := httpNet.ListenAndServe(addr, recoveryHandler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
