package handlers

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMessage(c *gin.Context) {
	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erorr": err.Error()})
		return
	}

	// Ensure that the message has both sender_id and text
	if message.SenderID == 0 || message.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SenderID and Text are required"})
		return
	}

	err := repositories.AddMessage(message.SenderID, message.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message added successfully", "content": message})
}

func GetUserMessages(c *gin.Context) {
	messages, err := repositories.GetUserMessages()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to retrieve messages for this user(?)"})
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
