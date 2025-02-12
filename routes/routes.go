package routes

import (
	"chat_app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Load HTML templates from the "templates" directory
	router.LoadHTMLGlob("templates/*")

	// Serve static files from /root/frontend
	router.Static("/frontend", "/root/frontend")

	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping) // Health check

		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)

		api.GET("/messages", handlers.GetUserMessages)
		api.POST("/messages", handlers.PostMessage)
	}

	// Define routes to render HTML pages
	router.GET("/", handlers.HomePage)
	router.GET("/about", handlers.AboutPage)
}
