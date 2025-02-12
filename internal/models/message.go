package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID uint   `json:"sender_id"`
	Text     string `json:"text"`
	Sender   User   `gorm:"foreignKey:SenderID"`
}
