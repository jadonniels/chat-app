package handlers

import (
	"chat_app/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users."})
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		"users": users,
	})
}

func AboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "About Us",
	})
}
