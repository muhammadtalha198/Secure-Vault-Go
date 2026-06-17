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
	"github.com/muhammadtalha198/secure-vault-api/internal/server"
	// "github.com/muhammadtalha198/secure-vault-api/looger"
)

func main() {
	// Initialize logger first (so we can log startup)
	// logger.Setup()

	// Load configuration
	// Fail fast if required env vars are missing
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create server with all dependencies wired
	srv := server.New(cfg)

	// Set up graceful shutdown
	// Listen for OS signals (SIGTERM, SIGINT)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Start server in a goroutine so it doesn't block
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
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
