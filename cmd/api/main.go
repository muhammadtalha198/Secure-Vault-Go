package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/muhammadtalha198/secure-vault-api/internal/config"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("config.Load() failed: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	slog.Info("config loaded successfully")
	slog.Debug("config loaded successfully")

	fmt.Printf("DatabaseURL: %s\n", config.DatabaseURL)
	fmt.Printf("RedisURL: %s\n", config.RedisURL)
	fmt.Printf("Port: %s\n", config.Port)
	fmt.Printf("CORSOrigin: %s\n", config.CORSOrigin)
	fmt.Printf("RateLimitLogin: %d\n", config.RateLimitLogin)
	fmt.Printf("StorageQuotaFreeGB: %d\n", config.StorageQuotaFreeGB)
	fmt.Printf("MaxUploadSizeMB: %d\n", config.MaxUploadSizeMB)
	fmt.Printf("JWT private key loaded: %t\n", config.JWTPrivateKey != nil)
	fmt.Printf("JWT public key loaded: %t\n", config.JWTPublicKey != nil)
}
