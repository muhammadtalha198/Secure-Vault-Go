package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadtalha198/secure-vault-api/internal/config"
	"github.com/muhammadtalha198/secure-vault-api/internal/handler"
)

// Server wraps the http server and its dependencies
type Server struct {
	router        *gin.Engine
	config        *config.Config
	httpServer    *http.Server
	healthHandler *handler.HealthHandler
}

// NewServer creates a new server with the given configuration
func New(cfg *config.Config) *Server {

	// Create Gin router
	// Use gin.New() for explicit control over middleware

	router := gin.New()

	// Create server
	server := &Server{
		router:        router,
		config:        cfg,
		healthHandler: handler.NewHealthHandler(),
	}

	// We'll add middleware and routes in separate functions
	server.registerRoutes()

	return server
}

// Start begins listening for HTTP requests
func (s *Server) Start() error {

	s.httpServer = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.router,
		// Timeouts prevent slow clients from holding connections
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
