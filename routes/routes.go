package routes

import (
	"chat_app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Load HTML templates from the "templates" directory
	router.LoadHTMLGlob("templates/*")

	// Serve static files (CSS, JS, images) from "public" directory
	router.Static("/public", "./public")

	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping) // Health check
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
	}

	// Define routes to render HTML pages
	router.GET("/", handlers.HomePage)
	router.GET("/about", handlers.AboutPage)
}
