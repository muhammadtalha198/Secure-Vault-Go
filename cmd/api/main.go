package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read PORT from environment; if not set, use 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Gin engine = your HTTP server + router
	router := gin.Default()

	// When someone visits GET /health, return JSON
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start listening on :3000 (or whatever PORT is)
	log.Printf("API listening on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
