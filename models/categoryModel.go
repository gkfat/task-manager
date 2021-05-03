package models

import (
	"time"

	"gorm.io/gorm"
)

// Category
type Category struct {
	gorm.Model
	ID        uint
	CreatedAt time.Time `gorm:"type:datetime;"`
	UpdatedAt time.Time `gorm:"type:datetime;"`
	DeletedAt time.Time `gorm:"type:datetime;"`
	Title     string    `gorm:"type:varchar(255)"`
}
