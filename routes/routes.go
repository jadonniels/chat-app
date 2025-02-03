package routes

import (
	"chat_app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping) // Health check
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
	}
}
