package server

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the HTTP routes for the server
func (s *Server) registerRoutes() {
	// Health check (no auth required)
	s.router.GET("/health", s.healthHandler.HandleHealth)

	// API v1 group
	v1 := s.router.Group("/api/v1")
	{
		// Auth routes (no auth required)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", s.handleNotImplemented)
			auth.POST("/login", s.handleNotImplemented)
			auth.POST("/refresh", s.handleNotImplemented)
		}

		// Protected routes (auth required - will add middleware later)
		protected := v1.Group("")
		{
			protected.GET("/documents", s.handleNotImplemented)
			protected.POST("/documents/upload", s.handleNotImplemented)
		}
	}
}

// handleNotImplemented returns 501 for routes not yet built
func (s *Server) handleNotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}
