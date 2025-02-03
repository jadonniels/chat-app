package repositories

import (
	"chat_app/config"
	"chat_app/internal/models"
)

// THIS PACKAGE IS RESPONSIBLE FOR DIRECTLY QUERYING THE DATABASE.
// THERE SHOULD BE NO BUSINESS LOGIC!

// GetAllUsers fetches all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

// CreateUser inserts a new user
func CreateUser(user models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}
