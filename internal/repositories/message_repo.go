package repositories

import (
	"chat_app/config"
	"chat_app/internal/models"
)

// THIS PACKAGE IS RESPONSIBLE ONLY FOR DIRECTLY QUERYING THE DATABASE.

func AddMessage(senderID uint, text string) error {
	message := models.Message{
		SenderID: senderID,
		Text:     text,
	}

	result := config.DB.Create(&message)
	return result.Error
}

// FIXME - Consider adding a "where user_id = ?" when going to chat room
func GetUserMessages() ([]models.Message, error) {
	var messages []models.Message

	// Preload the Sender (User) information along with the messages
	result := config.DB.Preload("Sender").Find(&messages)
	return messages, result.Error
}
