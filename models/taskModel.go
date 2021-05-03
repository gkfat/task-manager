package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// Task
type Task struct {
	gorm.Model
	ID          uint
	CreatedAt   time.Time `gorm:"type:datetime;"`
	UpdatedAt   time.Time `gorm:"type:datetime;"`
	DeletedAt   time.Time `gorm:"type:datetime;"`
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Category    uint
	Month       string `gorm:"type:varchar(255)"`
	Week        string `gorm:"type:varchar(255)"`
	Weekday     string `gorm:"type:varchar(255)"`
}
