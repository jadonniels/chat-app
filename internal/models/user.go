package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // includes the following fields by default: ID, CreatedAt, UpdatedAt, and DeletedAt
	Name       string `json:"name"`
}
