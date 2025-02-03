package handlers

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// THIS IS ALL OF THE API HANDLERS AND TIES ALL INTERNAL FUNCTIONALITY TOGETHER.

// check if server is running
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// GetUsers fetches users from the database
func GetUsers(c *gin.Context) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser saves a new user in the database
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
