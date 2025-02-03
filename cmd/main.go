package main

import (
	"log"
	"os"

	"chat_app/config"
	"chat_app/pkg/middleware"
	"chat_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.InitDB()

	// Load environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Gin router
	router := gin.Default()
	router.Use(middleware.Logger())

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	log.Println("Server running on port:", port)
	router.Run(":" + port)
}
