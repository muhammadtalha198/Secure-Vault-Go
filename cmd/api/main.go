package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muhammadtalha198/secure-vault-api/internal/config"
	"github.com/muhammadtalha198/secure-vault-api/internal/repository"
	"github.com/muhammadtalha198/secure-vault-api/internal/server"

	_ "github.com/muhammadtalha198/secure-vault-api/docs"
)

// @title           Secure Vault API
// @version         1.0
// @description     End-to-end encrypted document vault API
// @host            localhost:3000
// @BasePath        /
// @schemes         http
func main() {
	// Initialize logger first (so we can log startup)
	// logger.Setup()

	// Load configuration
	// Fail fast if required env vars are missing
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbPool, err := repository.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()
	log.Println("Database connected successfully")

	srv := server.New(cfg, dbPool)

	// Set up graceful shutdown
	// Listen for OS signals (SIGTERM, SIGINT)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Start server in a goroutine so it doesn't block
	go func() {
		baseURL := "http://localhost:" + cfg.Port
		log.Printf("Server running at %s", baseURL)
		log.Printf("Health check: %s/health", baseURL)
		log.Printf("DB health:    %s/health/db", baseURL)
		log.Printf("Swagger UI:   %s/swagger/index.html", baseURL)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Block until we receive a shutdown signal
	<-quit
	log.Println("Shutting down server...")

	// Give active requests 10 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
