package models

import (
    "gorm.io/gorm"
)

// Product represents the product model
type Product struct {
    gorm.Model
    UserID    string `gorm:"column:user_id"`
    ProductID string `gorm:"column:product_id"`
}
