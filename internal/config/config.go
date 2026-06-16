package config

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// Config holds all application configuration
// Every field comes from environment variables
type Config struct {
	// Database
	DatabaseURL string

	// Cache
	RedisURL string

	// JWT
	JWTPrivateKey *rsa.PrivateKey
	JWTPublicKey  *rsa.PublicKey

	// Server
	Port       string
	CORSOrigin string

	// Rate Limiting
	RateLimitLogin int

	// Storage
	StorageQuotaFreeGB int
	MaxUploadSizeMB    int
}

// Load reads environment variables and returns a validated Config
func Load() (*Config, error) {
	// Try to load local .env file for development.
	// If the file is missing, continue and rely on process environment variables.
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to read .env file: %w", err)
	}

	// Tell Viper to read environment variables (these override .env values).
	viper.AutomaticEnv()

	// Set defaults for optional values
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("RATE_LIMIT_LOGIN", 5)
	viper.SetDefault("STORAGE_QUOTA_FREE_GB", 1)
	viper.SetDefault("MAX_UPLOAD_SIZE_MB", 100)
	viper.SetDefault("JWT_PRIVATE_KEY_PATH", "secrets/jwt_private.pem")
	viper.SetDefault("JWT_PUBLIC_KEY_PATH", "secrets/jwt_public.pem")

	// Load required values
	cfg := &Config{
		DatabaseURL:        viper.GetString("DATABASE_URL"),
		RedisURL:           viper.GetString("REDIS_URL"),
		Port:               viper.GetString("PORT"),
		CORSOrigin:         viper.GetString("CORS_ORIGIN"),
		RateLimitLogin:     viper.GetInt("RATE_LIMIT_LOGIN"),
		StorageQuotaFreeGB: viper.GetInt("STORAGE_QUOTA_FREE_GB"),
		MaxUploadSizeMB:    viper.GetInt("MAX_UPLOAD_SIZE_MB"),
	}

	privateKeyPEM, err := loadKeyPEM("JWT_PRIVATE_KEY_PEM", "JWT_PRIVATE_KEY_PATH")
	if err != nil {
		return nil, err
	}

	publicKeyPEM, err := loadKeyPEM("JWT_PUBLIC_KEY_PEM", "JWT_PUBLIC_KEY_PATH")
	if err != nil {
		return nil, err
	}

	// Validate required fields
	required := map[string]string{
		"DATABASE_URL": cfg.DatabaseURL,
		"REDIS_URL":    cfg.RedisURL,
	}

	for name, value := range required {
		if value == "" {
			return nil, fmt.Errorf("required environment variable %s is not set", name)
		}
	}

	// Parse RSA private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT private key: %w", err)
	}
	cfg.JWTPrivateKey = privateKey

	// Parse RSA public key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT public key: %w", err)
	}
	cfg.JWTPublicKey = publicKey

	return cfg, nil
}

func loadKeyPEM(pemVarName, pathVarName string) ([]byte, error) {
	if keyPEM := viper.GetString(pemVarName); keyPEM != "" {
		return []byte(keyPEM), nil
	}

	keyPath := viper.GetString(pathVarName)
	if keyPath == "" {
		return nil, fmt.Errorf("required environment variable %s or %s is not set", pemVarName, pathVarName)
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key from %s (%s): %w", pathVarName, keyPath, err)
	}

	return keyPEM, nil
}
