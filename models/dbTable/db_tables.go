package dbTable

import (
	"time"

	"gorm.io/gorm"
)

// User
type User struct {
	gorm.Model
	ID        uint
	CreatedAt time.Time      `gorm:"type:date;"`
	UpdatedAt time.Time      `gorm:"type:date;"`
	DeletedAt gorm.DeletedAt `gorm:"type:date;"`
	UserName  string         `gorm:"type:varchar(255)"`
	Account   string         `gorm:"type:varchar(255)"`
	Password  string         `gorm:"not null"`
	Token     string         `gorm:"type:varchar(255)"`
}

// Category
type Category struct {
	gorm.Model
	ID        uint
	CreatedAt time.Time      `gorm:"type:date;"`
	UpdatedAt time.Time      `gorm:"type:date;"`
	DeletedAt gorm.DeletedAt `gorm:"type:date;"`
	Title     string         `gorm:"type:varchar(255)"`
}

// Task
type Task struct {
	gorm.Model
	ID          uint
	CreatedAt   time.Time      `gorm:"type:date;"`
	UpdatedAt   time.Time      `gorm:"type:date;"`
	DeletedAt   gorm.DeletedAt `gorm:"type:date;"`
	Title       string         `gorm:"type:varchar(255)"`
	Description string         `gorm:"type:varchar(255)"`
	CategoryID  uint
	LimitDate   string `gorm:"type:date;"`
}
