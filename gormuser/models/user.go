package models

import (
    "gorm.io/gorm"
)

// User represents the user model
type User struct {
    gorm.Model
    UserID string `gorm:"column:user_id"`
    Email  string
}
