package models

import "gorm.io/gorm"

// User
type User struct {
	gorm.Model
	ID       uint
	UserName string
	Account  string
	Password string `gorm:"not null"`
	Token    string
}
